package event

import (
	"reflect"

	"github.com/xozrc/eventsourcing/event"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

type CancelEvent struct {
	sourceId eventsourcingtypes.Guid
	version  int64
}

func (ie *CancelEvent) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *CancelEvent) Version() int64 {
	return ie.version
}

func NewCancelEvent(sourceId eventsourcingtypes.Guid, version int64) event.VersionedEvent {
	return &CancelEvent{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	cancelEventKey := reflect.TypeOf((*CancelEvent)(nil)).Name()
	event.RegisterVersionEventFactory(cancelEventKey, event.VersionEventFactoryFunc(NewCancelEvent))
}
