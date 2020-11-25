package main

import "github.com/multiformats/go-multiaddr"


type Config struct {
	NetworkId		 string
	BootstrapPeers   []multiaddr.Multiaddr
	ProtocolID       string
}



var DefaultConfig = & Config {
	NetworkId: 		"bc team",
	BootstrapPeers: []multiaddr.Multiaddr{
		//StringsToAddr("/ip4/192.168.1.6/tcp/52363/ipfs/QmYnXLw9u3zN4EcQhT89ZnFqpeEn5tsaXRXuzUkYDgbyvv"),
	},
	ProtocolID: 	"/bc_p2p/1.1.0",
}

func StringsToAddr(addrString string) multiaddr.Multiaddr {
	addr, _ := multiaddr.NewMultiaddr(addrString)
	return addr
}

