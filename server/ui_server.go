package server

import (
	"strings"

	c "github.com/Yangiboev/tcp-server-client/client"
	"github.com/Yangiboev/tcp-server-client/ui"
	"github.com/marcusolsson/tui-go"
)

func StartUIServer(s ServerI) {
	chatView := ui.NewChatView()
	ui, err := tui.New(chatView)
	if err != nil {
		panic(err)
	}
	quit := func() { ui.Quit() }
	ui.SetKeybinding("Esc", quit)
	ui.SetKeybinding("Ctrl+c", quit)
	chatView.OnSubmit(func(msg string) {
		strs := strings.Split(msg, " ")
		if len(strs) > 0 {
			if strs[0] == "tag" {
				go s.Send(strs[1], c.MessageCommand{
					Message: msg,
					Name:    "Server",
				})
			} else {
				go s.Broadcast(c.MessageCommand{
					Message: msg,
					Name:    "Server",
				})
			}
		}
	})
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
