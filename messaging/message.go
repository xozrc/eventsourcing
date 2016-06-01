package messaging

type Sender interface {
	Send([]byte) error
}

type Handler interface {
	Handle(msg []byte) error
}

type Receiver interface {
	Start(h Handler) error
	Stop() error

}
