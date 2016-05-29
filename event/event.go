package event

import "golang.org/x/net/context"

type Guid int64
type EventType int64

type Event interface {
	//from aggregate guid
	SourceId() Guid
}

type VersionedEvent interface {
	Event
	Version() int64
}


type EventSourceder interface{
	Guid() Guid
	Version() int
	SetVersion(int)
	Events() []VersionedEvent
	AddEvent(e VersionedEvent)
}

type EventSourced struct {
	EventSourceder
}

func(es *EventSourced) LoadFrom() error{
	for _,e := range es.Events(){
		es.SetVersion(e.Version())
	}
	
}

func(es *EventSourced) Update(e VersionedEvent){
	
}



type EventHandler interface {
	ApplyEvent(ctx context.Context, e Event) error
}

type EventPublisher interface{
	SendEvent(e Event) error
}

