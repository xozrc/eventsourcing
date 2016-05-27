package command

import (
	"errors"

	"golang.org/x/net/context"
)

var (
	EventHandlerNoFound = errors.New("event handler no found")
)

const (
	initEventHandlerSize = 10
)

type EventDispatcher interface {
	DispatchEvent(ctx context.Context, e Event) error
	Register(et EventType, ch EventHandler)
	Deregister(et EventType)
}

type eventDispatcher struct {
	handlersMap map[EventType]EventHandler
}

func (ed *eventDispatcher) DispatchEvent(ctx context.Context, e Event) error {
	h, ok := ed.handlersMap[c.EventType()]
	if !ok {
		return EventHandlerNoFound
	}
	return h.ApplyEvent(ctx, e)
}

func (cd *eventDispatcher) Register(et EventType, eh EventHandler) {
	cd.handlersMap[et] = eh
}

func (cd *eventDispatcher) Deregister(et EventType) {
	delete(cd.handlersMap, et)
}

func NewEventDispatcher() EventDispatcher {
	cd := &eventDispatcher{}
	cd.handlersMap = make(map[EventType]EventHandler, initEventHandlerSize)
	return cd
}
