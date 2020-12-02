package main

import (
	"bufio"
	"errors"
	mapset "github.com/deckarep/golang-set"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

var (
	errAlreadyRegistered = errors.New("peer is already registered")
	errNotRegistered     = errors.New("peer is not registered")
)


type Peer struct {

	stream 		network.Stream

	id			peer.ID

	rw			*bufio.ReadWriter

	knownMsg	mapset.Set                // Set of messages hashes known to be known by this peer

}


func NewPeer(stream network.Stream) *Peer {
	p := &Peer {
		stream: 	stream,
		id: 		stream.Conn().RemotePeer(),
		rw: 		bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream)),
		knownMsg:	mapset.NewSet(),
	}


	return p
}

func (p *Peer) ID() peer.ID  {
	return p.id
}

func (p *Peer) Close() error {
	return p.stream.Close()
}


// peerSet represents the collection of active peers currently participating in
// the messaging protocol
type peerSet struct {
	peers  map[peer.ID]*Peer
}

// newPeerSet creates a new peer set to track the active participants.
func newPeerSet() *peerSet {
	return &peerSet{
		peers: make(map[peer.ID]*Peer),
	}
}

// Register injects a new peer into the working set, or returns an error if the
// peer is already known.
func (ps *peerSet) Register(p *Peer) error {

	if _, ok := ps.peers[p.id]; ok {
		return errAlreadyRegistered
	}

	ps.peers[p.id] = p

	return nil
}

// Unregister removes a remote peer from the active set, disabling any further
// actions to/from that particular entity.
func (ps *peerSet) Unregister(id peer.ID) error {
	p, ok := ps.peers[id]
	if !ok {
		return errNotRegistered
	}

	delete(ps.peers, id)

	return p.Close()
}

func (ps *peerSet) Delete(id peer.ID)  {
	delete(ps.peers, id)
}

// Peer retrieves the registered peer with the given id.
func (ps *peerSet) Peer(id peer.ID) *Peer {
	return ps.peers[id]
}

// Len returns if the current number of peers in the set.
func (ps *peerSet) Len() int {
	return len(ps.peers)
}

// MarkMessage marks a message as known for the peer, ensuring that it
// will never be propagated to this particular peer.
func (p *Peer) MarkMessage(hash Hash) {
	p.knownMsg.Add(hash)
}

// PeersWithoutMsg retrieves a list of peers that do not have a given message
// in their set of known hashes.
func (ps *peerSet) PeersWithoutMsg(hash Hash) []*Peer {

	list := make([]*Peer, 0, len(ps.peers))

	for _, p := range ps.peers {
		if !p.knownMsg.Contains(hash) {
			list = append(list, p)
		}
	}

	return list
}