package event

import (
	"reflect"

	"github.com/xozrc/eventsourcing/event"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

type VerifyEvent struct {
	sourceId eventsourcingtypes.Guid
	version  int64
}

func (ie *VerifyEvent) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *VerifyEvent) Version() int64 {
	return ie.version
}

func NewVerifyEvent(sourceId eventsourcingtypes.Guid, version int64) event.VersionedEvent {
	return &VerifyEvent{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	verifyEventKey := reflect.TypeOf((*VerifyEvent)(nil)).Name()
	event.RegisterVersionEventFactory(verifyEventKey, event.VersionEventFactoryFunc(NewVerifyEvent))
}
