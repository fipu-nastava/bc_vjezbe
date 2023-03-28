package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"log"
	"strings"
)
type App struct {

	inputStream  chan string
	outputStream chan *Message

	messages  	 []*Message
}

func NewApp(inputStream chan string, outputStream chan *Message) *App  {
	app := &App{
		inputStream:	inputStream,
		outputStream:	outputStream,
	}

	return app
}

func (a *App) sendMsg(text string)  {

	a.messages = append(a.messages, &Message{Sender: "You", Msg: text})

	a.inputStream <- text
}


func (a *App) readMessages()  {
	for {
		m := <- a.outputStream
		a.messages = append(a.messages, m)
	}
}

func (a *App) loop(w *app.Window) error {

	th := material.NewTheme()
	gtx := &layout.Context{
		Queue: w.Queue(),
	}


	history	   := &layout.List{
		Axis: 			layout.Vertical,
		ScrollToEnd: 	true,
		Alignment: 		layout.Start,
	}
	lineEditor := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	list       := &layout.List{
		Axis: layout.Vertical,
	}


	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			a.draw(gtx, th, history, lineEditor, list)
			e.Frame(gtx.Ops)
		}
	}
}

func (a *App) draw(gtx *layout.Context, th *material.Theme, history *layout.List, lineEditor *widget.Editor, list *layout.List) {
	widgets := []func(){

		func() {

			gtx.Constraints.Height.Max = gtx.Px(unit.Dp(700))
			gtx.Constraints.Width.Max = gtx.Px(unit.Dp(750))

			history.Layout(gtx, len(a.messages), func(index int) {

				m := a.messages[index]
				l := th.Label(unit.Dp(18), fmt.Sprintf("%s> %s", m.Sender, m.Msg))
				l.Layout(gtx)

			})

		},

		func() {
			e := th.Editor("Write message")
			e.Font.Style = text.Italic
			e.Layout(gtx, lineEditor)
			for _, e := range lineEditor.Events(gtx) {
				if e, ok := e.(widget.SubmitEvent); ok {

					txt := e.Text
					lineEditor.SetText("")
					if strings.TrimSpace(txt) == "" {
						return
					}
					a.sendMsg(txt)
				}
			}
		},
	}


	list.Layout(gtx, len(widgets), func(i int) {
		layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})

}


func (a *App) Start() {

	go a.readMessages()

	gofont.Register()
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(800), unit.Dp(800)),
			app.Title("P2P Chat"),
		)
		if err := a.loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}