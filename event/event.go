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

type EventHandler interface {
	ApplyEvent(ctx context.Context, e Event) error
}


type EventSourced interface {
	Guid() Guid
	Version() int
	Events() []Event
	LoadFrom(pes []Event) error
}

type EventStore interface {
	Save(partitionKey string, es []Event) error
}

type EventStoreBusPublisher interface {
	SendAsync(partitionKey string) error
}