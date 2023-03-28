package main

import "github.com/multiformats/go-multiaddr"


type Config struct {
	NodeName		 string
	NetworkId		 string
	BootstrapPeers   []multiaddr.Multiaddr
	ProtocolID       string
	MaxPeers 	     int
}



var DefaultConfig = & Config {
	NodeName: 		"Robert",
	NetworkId: 		"bc team",
	BootstrapPeers: []multiaddr.Multiaddr{
		StringsToAddr("/ip4/192.168.43.30/tcp/60090/ipfs/QmdgrpPVktK6a4pYtf9hUTWNUN8VhfTRqAYm6mTNqdjL5E"),
	},
	ProtocolID: 	"/bc_p2p/1.1.0",
	MaxPeers:  		25,
}

func StringsToAddr(addrString string) multiaddr.Multiaddr {
	addr, _ := multiaddr.NewMultiaddr(addrString)
	return addr
}

