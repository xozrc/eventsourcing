package event

import (
	"reflect"

	"github.com/xozrc/eventsourcing/event"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

type InitEvent struct {
	sourceId eventsourcingtypes.Guid
	version  int64
}

func (ie *InitEvent) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *InitEvent) Version() int64 {
	return ie.version
}

func NewInitEvent(sourceId eventsourcingtypes.Guid, version int64) event.VersionedEvent {
	return &InitEvent{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	initEventKey := reflect.TypeOf((*InitEvent)(nil)).Name()
	event.RegisterVersionEventFactory(initEventKey, event.VersionEventFactoryFunc(NewInitEvent))
}
