package repository

import (
	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/types"
)

type Repository interface {
	Find(id types.Guid, es *event.EventSourced) (err error)
}
