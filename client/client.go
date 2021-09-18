package client

type ClientI interface {
	Start()
	Close()
	Dial(address string) error
	Send(command interface{}) error
	SetName(name string) error
	SendMessage(message string) error
	Incoming() chan MessageCommand
	Error() chan error
}
