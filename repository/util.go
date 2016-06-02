package repository

import (
	"fmt"

	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/store"
	"github.com/xozrc/eventsourcing/types"
)

func GetPartitionKey(sourceType string, id types.Guid) string {
	return fmt.Sprintf("%s_%d", sourceType, id)
}

func ConvertEventToData(e event.VersionedEvent) (ed *store.EventData, err error) {
	ed := &store.EventData{}
	ed.SourceId = fmt.Sprintf("%d", e.SourceId())
	return
}

func ConvertDataToEvent(ed *store.EventData) (e event.VersionedEvent, err error) {
	return
}

func snapShotEventSourced(es *event.EventSourced) (bs []byte, err error) {
	return
}

type Marshaller interface {
	Marshal(e *event.VersionedEvent) (*EventData, error)
}

type Unmarshaller interface {
	Unmarshal(data *EventData) (*event.VersionedEvent, error)
}
