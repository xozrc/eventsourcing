package rabbitmq

import (
	"os"

	"github.com/streadway/amqp"
)

const (
	//amqp://guest:guest@127.0.0.1:5672/
	amqpUrl = "AMQP_URL"
)

type RabbitMQSender struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   *amqp.Queue
}

func (s *RabbitMQSender) Send(body []byte) error {
	return s.channel.Publish("", s.queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
}

func (s *RabbitMQSender) Close() error {
	return s.conn.Close()
}

func New(au string, qn string) (s *RabbitMQSender, err error) {

	s = &RabbitMQSender{}

	if au == "" {
		au = os.Getenv("AMQP_URL")
	}

	s.conn, err = amqp.Dial(au)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			s.conn.Close()
		}
	}()

	s.channel, err = s.conn.Channel()
	if err != nil {
		return
	}

	qu, err := s.channel.QueueDeclare(qn, true, false, true, false, nil)
	if err != nil {
		return
	}
	s.queue = &qu

	return
}
