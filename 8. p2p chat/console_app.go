package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleApp struct {

	inputStream  chan string
	outputStream chan *Message

}

func NewConsoleApp(inputStream chan string, outputStream chan *Message) *ConsoleApp  {
	app := &ConsoleApp{
		inputStream:	inputStream,
		outputStream:	outputStream,
	}

	return app
}

func (a *ConsoleApp) sendMsg(text string)  {

	showMsg(&Message{Sender: "You", Msg: text})

	a.inputStream <- text
}

func (a *ConsoleApp) readInput()  {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		if text = strings.Trim(text, "\n"); text == "" {
			continue
		}
		a.sendMsg(text)
	}


}


func (a *ConsoleApp) readMessages()  {
	for {
		m := <- a.outputStream
		showMsg(m)
	}
}

func showMsg(m *Message)  {
	fmt.Println(m.Sender, ">", m.Msg)
}

func (a *ConsoleApp) Start() {

	go a.readMessages()
	go a.readInput()
}