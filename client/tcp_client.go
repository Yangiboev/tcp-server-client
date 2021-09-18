package client

import (
	"log"
	"net"
	"strings"
)

type TcpChatClient struct {
	conn      net.Conn
	cmdReader *CommandReader
	cmdWriter *CommandWriter
	name      string
	error     chan error
	incoming  chan MessageCommand
}

func NewClient() ClientI {
	return &TcpChatClient{
		incoming: make(chan MessageCommand),
		error:    make(chan error),
	}
}

func (c *TcpChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
		c.cmdReader = NewCommandReader(conn)
		c.cmdWriter = NewCommandWriter(conn)
	}

	return err
}

func (c *TcpChatClient) Start() {
	for {
		cmd, err := c.cmdReader.Read()

		if err != nil {
			c.error <- err
			return
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case MessageCommand:
				c.incoming <- v
			default:
				log.Printf("Unknown command: %v", v)
			}
		}
	}
}

func (c *TcpChatClient) Close() {
	c.conn.Close()
}

func (c *TcpChatClient) Incoming() chan MessageCommand {
	return c.incoming
}

func (c *TcpChatClient) Error() chan error {
	return c.error
}

func (c *TcpChatClient) Send(command interface{}) error {
	return c.cmdWriter.Write(command)
}

func (c *TcpChatClient) SetName(name string) error {
	return c.Send(NameCommand{Name: name})
}

func (c *TcpChatClient) SendMessage(message string) error {
	strs := strings.Split(message, " ")
	if strs[0] == "tag" {
		if strs[1] != "" {
			return c.Send(MessageCommand{
				Message: message,
				Name:    strs[1],
			})
		}
	}
	return c.Send(SendCommand{
		Message: message,
	})
}
