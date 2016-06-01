package event

import "github.com/xozrc/eventsourcing/types"

//Event ...
type Event interface {
	SourceId() types.Guid
}

//VersionedEvent ...
type VersionedEvent interface {
	Event
	Version() int64
}

type EventSourcedModel interface {
	Id() Guid
	Version() int64
	ApplyEvent(ve VersionedEvent) error
	Events() []VersionedEvent
	Payload() []byte
}

type EventSourced struct {
	EventSourcedModel
}

func (esb *EventSourced) LoadFrom(pes []VersionedEvent) error {
	for _, e := range pes {
		err := esb.ApplyEvent(e)
		if err != nil {
			return err
		}
	}
	return nil
}
