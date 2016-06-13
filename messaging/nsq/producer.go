package nsq

import (
	"github.com/nsqio/go-nsq"
)

type NSQProducer struct {
	producer *nsq.Producer
	topic    string
}

func (nsqp *NSQProducer) Publish(msg []byte) error {
	return nsqp.producer.Publish(nsqp.topic, msg)
}

func (nsqp *NSQProducer) PublishMulti(msgs [][]byte) error {
	for _, msg := range msgs {
		err := nsqp.Publish(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (nsqd *NSQProducer) Close() error {
	nsqd.producer.Stop()
	return nil
}

func NewProducer(p *nsq.Producer, topic string) (nsqp *NSQProducer, err error) {
	nsqp = &NSQProducer{}
	nsqp.producer = p
	nsqp.topic = topic
	return
}
