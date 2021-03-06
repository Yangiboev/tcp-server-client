package ui

import (
	"fmt"
	"io"
	"time"

	"github.com/Yangiboev/tcp-server-client/client"
	"github.com/marcusolsson/tui-go"
)

func StartUi(c client.ClientI) {
	loginView := NewLoginView()
	chatView := NewChatView()

	ui, err := tui.New(loginView)
	if err != nil {
		panic(err)
	}

	quitChan := make(chan struct{})

	quit := func() {
		ui.Quit()
		quitChan <- struct{}{}
	}

	ui.SetKeybinding("Esc", quit)
	ui.SetKeybinding("Ctrl+c", quit)

	loginView.OnLogin(func(username string) {
		c.SetName(username)
		ui.SetWidget(chatView)
	})

	chatView.OnSubmit(func(msg string) {
		c.SendMessage(msg)
	})

	go chattingProcess(c, ui, chatView, quitChan)

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

func chattingProcess(c client.ClientI, ui tui.UI, chatView *ChatView, quitChan chan struct{}) {
	for {
		select {
		case err := <-c.Error():
			if err == io.EOF {
				ui.Update(func() {
					chatView.AddMessage("Connection closed connection from server.")
				})
			} else {
				panic(err)
			}
		case msg := <-c.Incoming():
			ui.Update(func() {
				chatView.AddMessage(fmt.Sprintf("%s-%s: %s", time.Now().Format("2006.01.02 15:04:05:"), msg.Name, msg.Message))
			})
		case <-quitChan:
			return
		}
	}
}
