package nsq_test

import "testing"

import (
	"github.com/nsqio/go-nsq"
	. "github.com/xozrc/mq/mq/nsq"
)

const (
	testTopic = "test"
	testUrl   = "192.168.99.109:4150"
)

func setUpProducer() (p *NSQProducer, err error) {
	tp, err := nsq.NewProducer(testUrl, nsq.NewConfig())
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tp.Stop()
		}
	}()

	p, err = NewProducer(tp, testTopic)
	if err != nil {
		return
	}
	return
}

func TestProducer(t *testing.T) {

	p, err := setUpProducer()
	if err != nil {
		t.Errorf("set up producer error:%s", err.Error())
	}

	defer p.Close()

	err = p.Publish([]byte("hello world"))
	if err != nil {
		t.Errorf("publish error :%s", err.Error())
	}
}

func BenchmarkProduce(b *testing.B) {
	p, err := setUpProducer()
	if err != nil {
		b.Fatalf("set up producer error :%s", err.Error())
	}
	defer p.Close()
	for i := 0; i < b.N; i++ {
		err := p.Publish([]byte("hello world"))
		if err != nil {
			b.Fatalf("publish error :%s", err.Error())
		}
	}
}
