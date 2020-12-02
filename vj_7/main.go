package main

import "github.com/ipfs/go-log"


func main() {


	log.SetLogLevel("discovery", "debug")
	log.SetLogLevel("node", "debug")



	inputStream  := make(chan string, 10)
	outputStream := make(chan *Message, 10)

	n := NewNode(DefaultConfig, inputStream, outputStream)
	n.Start()

	app := NewConsoleApp(inputStream, outputStream)
	app.Start()

}
