package v6

import (
	"context"
	"github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/network"
)

var nLogger = log.Logger("node")


type Node struct {
	*Config
	context.Context

	discovery   *discovery

}

func NewNode(c *Config) *Node {
	n := &Node {
		Config:  		c,
		Context: 		context.Background(),
	}

	n.discovery = NewDiscovery(n, n.Config)

	return n
}



func (n *Node) handleNewPeer(stream network.Stream)  {
	nLogger.Info("Found new pper ", stream.Conn().RemotePeer())
}

func (n *Node) startDiscovery()  {
	n.discovery.Start(n.handleNewPeer)
}


func (n *Node) Start()  {
	go n.startDiscovery()
}

