package rabbitmq_test

import (
	"testing"
)

import (
	. "github.com/xozrc/eventsourcing/messaging/rabbitmq"
)

func TestSend(t *testing.T) {
	sender, err := New("", "test")
	if err != nil {
		t.Fatalf("new error :%s", err.Error())
	}
	defer sender.Close()
	err = sender.Send([]byte("hello world"))
	if err != nil {
		t.Fatalf("send error :%s", err.Error())
	}
}
