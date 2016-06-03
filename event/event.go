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

//trivival versioned event
type TrivialVersionedEvent struct {
	sourceId types.Guid
	version  int64
}

func (tve *TrivialVersionedEvent) SourceId() types.Guid {
	return tve.sourceId
}

func (tve *TrivialVersionedEvent) Version() int64 {
	return tve.version
}

type VersionEventFactory interface {
	NewVersionEvent(sourceId types.Guid, version int64) VersionedEvent
}

func NewVersionEvent(sourceId types.Guid, version int64) VersionedEvent {
	return &TrivialVersionedEvent{
		sourceId: sourceId,
		version:  version,
	}
}

type EventSourced interface {
	SourceId() types.Guid
	Version() int64
	ApplyEvent(ve VersionedEvent) error
	Events() []VersionedEvent
	Payload() []byte
}
