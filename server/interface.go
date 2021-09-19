package server

type ServerI interface {
	Listen(address string) error
	Broadcast(command interface{}) error
	Send(name string, command interface{}) error
	Start()
	Close()
}
