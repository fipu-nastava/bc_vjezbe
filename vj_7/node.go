package main

import (
	"context"
	mapset "github.com/deckarep/golang-set"
	"github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/network"
)

var nLogger = log.Logger("node")


type Node struct {
	*Config
	context.Context

	discovery   *discovery

	peers 		*peerSet

	knownMsg	mapset.Set                // Set of messages hashes known to be known by this node

	counter     int

	inputStream  chan string
	outputStream chan *Message


}

func NewNode(c *Config, inputStream chan string, outputStream chan *Message) *Node {
	n := &Node {
		Config:  		c,
		Context: 		context.Background(),
		peers:   		newPeerSet(),
		knownMsg: 		mapset.NewSet(),
		counter: 		1,
		inputStream: 	inputStream,
		outputStream:	outputStream,
	}

	n.discovery = NewDiscovery(n, n.Config)

	return n
}

func (n *Node) markAsKnown(m *Message) bool  {
	known :=  n.knownMsg.Contains(m.Hash())

	if !known {
		n.knownMsg.Add(m.Hash())
	}

	return known
}


func (n *Node) handleNewPeer(stream network.Stream)  {

	if n.peers.Len() < n.MaxPeers {


		peer := NewPeer(stream)
		err := n.peers.Register(peer)

		if err != nil {
			nLogger.Warning("Failed to connect to peer: ", err)
			peer.Close()
			return
		}

		nLogger.Info("Connected to new peer: ", peer.ID())
		go n.readData(peer)

	} else {
		stream.Close()
		nLogger.Debug("MaxPeers reached, rejecting peer")
	}
}

func (n *Node) startDiscovery()  {
	n.discovery.Start(n.handleNewPeer)
}


func (n *Node) Start()  {
	go n.startDiscovery()
	go n.writeMessages()
}


func (n *Node) writeMessages() {

	for {

		data := n.readFromInput()

		n.counter += 1
		m := &Message{
			Sender: n.NodeName,
			Msg:	data,
			Nonce:	n.counter,
		}

		n.Broadcast(m)
	}
}


func (n *Node) readData(peer *Peer) {
	rw := peer.rw
	for {

		m, err := Deserialize(rw)


		if err != nil {
			n.peers.Unregister(peer.ID())
			nLogger.Warning("Reading from peer failed, removing peer ", peer.ID())
			return
		}

		// mark as know to peer as he send it
		peer.MarkMessage(m.Hash())

		if !n.markAsKnown(m) {
			// send to output if i haven't seen the message before
			n.sendToOutput(m)
		}


		n.Broadcast(m)
	}
}

func (n *Node) Broadcast(m *Message)  {

	nLogger.Debug("Broadcasting message: ", m)

	data := m.Serialize()

	for _, peer := range n.peers.PeersWithoutMsg(m.Hash()) {
		rw := peer.rw

		_, err := rw.Write(data)

		if err == nil {
			err = rw.Flush()
		}

		if err != nil {
			n.peers.Unregister(peer.ID())
			nLogger.Warning("Writing to peer failed, removing peer ", peer.ID())
			return
		}

		// mark message as known to peer
		peer.MarkMessage(m.Hash())
	}
}

func (n *Node) sendToOutput(m *Message)  {
	n.outputStream <- m
}

func (n *Node) readFromInput() string {
	data := <- n.inputStream
	return data
}