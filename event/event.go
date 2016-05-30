package event

import "golang.org/x/net/context"

//Event ...
type Event interface {
	SourceId() Guid
}

//VersionedEvent ...
type VersionedEvent interface {
	Event
	Version() int64
}


type EventPublisher interface{
	Events []Event
}

type EventSourced interface{
	Guid() Guid
	Version() int
	Events() []VersionedEvent
}

type EventSourcedBasic struct {
	id Guid
	version int64
	pendingEvents []VersionedEvent

}

func(esb *EventSourcedBasic) Id() Guid{
	return esb.id
}

func(esb *EventSourcedBasic) Version() int64{
	return esb.version
}

func(esb *EventSourcedBasic) Events() []VersionedEvent{
	return esb.pendingEvents
}

func(esb *EventSourcedBasic) LoadFrom(	pes []VersionedEvent) error {
	return nil	
}

func(esb *EventSourcedBasic) Update(e VersionedEvent) error {
  	e.SourceId = esl.Id;
	e.Version = esl.version + 1;
	esl.es.Handler
	esb.version = e.Version;
	esb.pendingEvents.Add(e);
	return nil	
}

func NewEventSourcedBasic(id Guid,pendingEvents []VersionedEvent) *EventSourcedBasic{
	esb := &EventSourcedBasic{}
	esb.id = id
	esb.pendingEvents = pendingEvents
	return esb
}






