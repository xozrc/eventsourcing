package event

import (
	"encoding/json"
)

import (
	"golang.org/x/net/context"
)

type EventProcessor struct {
	ed EventDispatcher
}

func (ep *EventProcessor) Handle(msg []byte) error {
	ed := &EventData{}
	err := json.Unmarshal(msg, ed)
	if err != nil {
		return err
	}

	et := ed.EventType

	//todo: time out context
	ctx := context.Background()

	return nil
}

func NewEventProcessor(ed EventDispatcher) *EventProcessor {
	ep := &EventProcessor{}
	ep.ed = ed
	return ep
}
