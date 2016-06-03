package repository

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/store"
	"github.com/xozrc/eventsourcing/types"
)

func GetPartitionKey(sourceType string, id types.Guid) string {
	return fmt.Sprintf("%s_%d", sourceType, id)
}

func ConvertEventToData(e event.VersionedEvent) (ed *store.EventData, err error) {
	ed = &store.EventData{}
	ed.SourceId = fmt.Sprintf("%d", e.SourceId())
	return
}

func ConvertDataToEvent(ed *store.EventData) (e event.VersionedEvent, err error) {
	return
}

func snapShotEventSourced(es *event.EventSourced) (bs []byte, err error) {
	return
}

func toData(e event.VersionedEvent) (*store.EventData, error) {
	ed := &store.EventData{}
	ed.PartitionKey = ""

	//json endcode event
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	ed.Payload = string(payload)
	ed.SourceId = fmt.Sprintf("%d", e.SourceId())
	ed.SourceType = reflect.TypeOf(e).Name()
	ed.Version = e.Version()

	return ed, nil
}
