package main

import "github.com/ipfs/go-log"


func main() {


	log.SetLogLevel("discovery", "debug")
	log.SetLogLevel("node", "debug")


	is := make(chan string, 10)
	os := make(chan *Message, 10)

	n := NewNode(DefaultConfig, is, os)
	n.Start()

	app := NewConsoleApp(is, os)
	app.Start()

	select {}
}
