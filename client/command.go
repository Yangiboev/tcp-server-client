package client

import "errors"

var (
	ErrC = errors.New("invalid command")
)

// Sending new message from client
type SendCommand struct {
	Message string
}

// Setting client display name
type NameCommand struct {
	Name string
}

// Notifying new messages
type MessageCommand struct {
	Name    string
	Message string
}
