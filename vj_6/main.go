package v6

import "github.com/ipfs/go-log"


func main() {


	log.SetLogLevel("discovery", "debug")
	log.SetLogLevel("node", "debug")



	n := NewNode(DefaultConfig)
	n.Start()
	select {}
}
