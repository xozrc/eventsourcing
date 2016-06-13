package event

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
	DispatchEvent(ctx context.Context, eventType string, e Event) error
	Register(et string, ch EventHandler)
}

type eventDispatcher struct {
	handlersMap map[string]EventHandler
}

func (ed *eventDispatcher) DispatchEvent(ctx context.Context, eventType string, e Event) error {
	h, ok := ed.handlersMap[eventType]
	if !ok {
		return EventHandlerNoFound
	}
	return h.HandleEvent(ctx, e)
}

func (cd *eventDispatcher) Register(et string, eh EventHandler) {
	cd.handlersMap[et] = eh
}

func NewEventDispatcher() EventDispatcher {
	cd := &eventDispatcher{}
	cd.handlersMap = make(map[string]EventHandler, initEventHandlerSize)
	return cd
}
