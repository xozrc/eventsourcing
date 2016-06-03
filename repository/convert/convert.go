package convert

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/store"
)

type Converter interface {
	ToData(e event.VersionedEvent) (*store.EventData, error)
	ToEvent(data *store.EventData) (event.VersionedEvent, error)
}

type defautlConverter struct {
}

func (dc *defautlConverter) ToData(e event.VersionedEvent) (*store.EventData, error) {
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

func (dc *defautlConverter) ToEvent(payload []byte, e event.VersionedEvent) error {
	return json.Unmarshal(payload, e)
}
