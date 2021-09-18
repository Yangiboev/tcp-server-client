package server

type ServerI interface {
	Listen(address string) error
	Broadcast(command interface{}) error
	Start()
	Close()
}
