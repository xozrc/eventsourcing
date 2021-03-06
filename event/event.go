package event

import (
	"reflect"

	"github.com/xozrc/eventsourcing/types"
)

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

type VersionEventFactoryFunc func(sourceId types.Guid, version int64) VersionedEvent

func (veff VersionEventFactoryFunc) NewVersionEvent(sourceId types.Guid, version int64) VersionedEvent {
	return veff(sourceId, version)
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

var (
	versionEventFactoryMap map[string]VersionEventFactory
)

func init() {
	versionEventFactoryMap = make(map[string]VersionEventFactory)

	triviKey := reflect.TypeOf((*TrivialVersionedEvent)(nil)).Name()
	RegisterVersionEventFactory(triviKey, VersionEventFactoryFunc(NewVersionEvent))
}

func RegisterVersionEventFactory(key string, vef VersionEventFactory) {
	versionEventFactoryMap[key] = vef
}

func GetVersionEventFactory(key string) VersionEventFactory {
	return versionEventFactoryMap[key]
}

type EventData struct {
	EventType string `json:"event_type"`
	Payload   []byte `json:"payload"`
}
