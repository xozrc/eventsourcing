package messaging

type Sender interface {
	Send([]byte) error
}

type Receiver interface {
	Start(h MessageHandler) error
	Stop() error
}

type MessageHandler interface {
	Handle(msg []byte) error
}
