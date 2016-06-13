package eventsourcing

import (
	orderevent "github.com/xozrc/eventsourcing/examples/order/eventsourcing/event"
)
import (
	"github.com/xozrc/eventsourcing/event"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

type OrderEventSourced struct {
	id      eventsourcingtypes.Guid
	version int64
	events  []event.VersionedEvent
	order   *Order
}

func (oes *OrderEventSourced) SourceId() eventsourcingtypes.Guid {
	return oes.id
}

func (oes *OrderEventSourced) Version() int64 {
	return oes.version
}

func (oes *OrderEventSourced) Events() []event.VersionedEvent {
	return oes.events
}

func (oes *OrderEventSourced) ApplyEvent(ve event.VersionedEvent) error {
	return nil
}

func (oes *OrderEventSourced) Payload() []byte {
	return nil
}

func (oes *OrderEventSourced) Update(e event.VersionedEvent) {
	oes.ApplyEvent(e)
	oes.events = append(oes.events, e)
	oes.version = e.Version()
}

func (oes *OrderEventSourced) Cancel() error {
	curVer := oes.Version() + 1
	e := orderevent.NewCancelEvent(oes.SourceId(), curVer)
	oes.Update(e)
	return nil
}

func (oes *OrderEventSourced) Confirm() error {
	curVer := oes.Version() + 1
	e := orderevent.NewConfirmEvent(oes.SourceId(), curVer)
	oes.Update(e)
	return nil
}

func (oes *OrderEventSourced) Pending() error {
	curVer := oes.Version() + 1
	e := orderevent.NewPendingEvent(oes.SourceId(), curVer)
	oes.Update(e)
	return nil
}

func (oes *OrderEventSourced) Verify() error {
	curVer := oes.Version() + 1
	e := orderevent.NewVerifyEvent(oes.SourceId(), curVer)
	oes.Update(e)
	return nil
}
