package repository

import (
	"encoding/json"
	"fmt"
	"reflect"
)

import (
	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/store"
	"github.com/xozrc/eventsourcing/types"
)

func GetPartitionKey(sourceType string, id types.Guid) string {
	return fmt.Sprintf("%s_%d", sourceType, id)
}

func snapShotEventSourced(es *event.EventSourced) (bs []byte, err error) {
	return
}

func ToData(st string, partitionKey string, e event.VersionedEvent) (*store.EventEntity, error) {
	ed := &store.EventEntity{}
	ed.PartitionKey = partitionKey
	ed.SourceType = st
	ed.EventType = reflect.TypeOf(e).Name()
	//json endcode event
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	ed.Payload = string(payload)
	ed.SourceId = fmt.Sprintf("%d", e.SourceId())
	ed.Version = e.Version()
	return ed, nil
}
