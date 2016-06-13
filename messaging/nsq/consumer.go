package nsq

import "github.com/nsqio/go-nsq"

type NSQConsumer struct {
	consumer    *nsq.Consumer
	addr        string
	receiveChan chan []byte
}

func (nsqc *NSQConsumer) Consume() (rc <-chan []byte, err error) {

	err = nsqc.consumer.ConnectToNSQLookupd(nsqc.addr)
	if err != nil {
		return
	}

	return nsqc.receiveChan, nil
}

func (nsqc *NSQConsumer) HandleMessage(msg *nsq.Message) error {
	nsqc.receiveChan <- msg.Body
	return nil
}

func NewConsumer(c *nsq.Consumer, addr string) (nsqc *NSQConsumer, err error) {
	nsqc = &NSQConsumer{}
	nsqc.consumer = c
	nsqc.addr = addr
	nsqc.consumer.AddHandler(nsqc)
	return
}
