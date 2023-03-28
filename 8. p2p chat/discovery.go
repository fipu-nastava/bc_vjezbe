package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	p2pdisc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	"sync"
	"time"

	"github.com/ipfs/go-log"
)

var dLogger = log.Logger("discovery")



type discovery struct {
	config			*Config
	ctx 			context.Context

	host		host.Host

	kDHT			*dht.IpfsDHT

}

func NewDiscovery(ctx context.Context, config *Config) *discovery {
	d := &discovery{
		ctx: 	ctx,
		config: config,
	}

	return d
}


func (d *discovery) createHost(newPeerStream func(stream network.Stream))  {

	// 0.0.0.0 will listen on any interface device.
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", 0))


	// libp2p.New constructs a new libp2p Host. Other options can be added here.
	host, err := libp2p.New(d.ctx,
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.NATPortMap(),
	)
	if err != nil {
		panic(err)
	}


	// Set a function as stream handler. This function is called when a peer
	// initiates a connection and starts a stream with this peer.
	host.SetStreamHandler(protocol.ID(d.config.ProtocolID), newPeerStream)


	dLogger.Info("Host created. We are:", host.ID())
	dLogger.Info(host.Addrs())

	for _, v := range host.Addrs() {
		dLogger.Info(v.String() + "/ipfs/" + host.ID().String())
	}

	d.host = host

}


func (d *discovery) createDHT()  {

	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	kademliaDHT, err := dht.New(d.ctx, d.host)
	if err != nil {
		panic(err)
	}

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table periodically.
	dLogger.Debug("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(d.ctx); err != nil {
		panic(err)
	}

	d.kDHT = kademliaDHT

}

func (d *discovery) connectBootstrapPeers()  {

	bootstrapPeers := d.config.BootstrapPeers
	if len(bootstrapPeers) == 0 {
		bootstrapPeers = dht.DefaultBootstrapPeers[2:3]
	}

	host := d.kDHT.Host()

	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	var wg sync.WaitGroup
	for _, peerAddr := range bootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)

		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := host.Connect(d.ctx, *peerinfo); err != nil {
				dLogger.Warning(err)
			} else {
				dLogger.Info("Connection established with bootstrap node:", *peerinfo)
			}
		}()
	}
	wg.Wait()


}

func (d *discovery) startDiscovery() <-chan peer.AddrInfo {

	networkId := d.config.NetworkId

	// We use a NetworkId point to announce our location.
	// This is like telling your friends to meet you at the Eiffel Tower.
	dLogger.Info("Announcing ourselves...")

	routingDiscovery := p2pdisc.NewRoutingDiscovery(d.kDHT)
	p2pdisc.Advertise(d.ctx, routingDiscovery, networkId)

	dLogger.Debug("Successfully announced!")

	// Now, look for others who have announced
	// This is like your friend telling you the location to meet you.
	dLogger.Debug("Searching for other peers...")

	peerChan, err := routingDiscovery.FindPeers(d.ctx, networkId)
	if err != nil {
		panic(err)
	}

	return peerChan
}

func (d *discovery) connectToPeers(	peerChan <-chan peer.AddrInfo, newPeerStream func(stream network.Stream))  {

	protocolId := protocol.ID(d.config.ProtocolID)
	host := d.kDHT.Host()

	for peer := range peerChan {
		// don't connect to yourself, that would be irrational
		if peer.ID == host.ID() {
			continue
		}

		dLogger.Debug("Connecting to: ", peer)
		// start a connection stream
		stream, err := host.NewStream(d.ctx, peer.ID, protocolId)

		if err != nil {
			dLogger.Debug("Connection failed: ", err)
			// if error continue to next peer
			continue
		}

		// if no errors invoke function
		newPeerStream(stream)

	}

	dLogger.Debug("Finished connecting to peers")
}


func (d *discovery) Start(newPeerStream func(stream network.Stream))  {

	d.createHost(newPeerStream)

	d.createDHT()

	d.connectBootstrapPeers()

	for {

		peerChan := d.startDiscovery()

		d.connectToPeers(peerChan, newPeerStream)

		time.Sleep(2 * time.Minute)
	}
}

