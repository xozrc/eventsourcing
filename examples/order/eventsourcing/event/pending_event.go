package event

import (
	"reflect"

	"github.com/xozrc/eventsourcing/event"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

type PendingEvent struct {
	sourceId eventsourcingtypes.Guid
	version  int64
}

func (ie *PendingEvent) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *PendingEvent) Version() int64 {
	return ie.version
}

func NewPendingEvent(sourceId eventsourcingtypes.Guid, version int64) event.VersionedEvent {
	return &PendingEvent{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	pendingEventKey := reflect.TypeOf((*PendingEvent)(nil)).Name()
	event.RegisterVersionEventFactory(pendingEventKey, event.VersionEventFactoryFunc(NewPendingEvent))
}
