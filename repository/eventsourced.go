package repository

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/xozrc/eventsourcing/cache"
	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/store"
)

type EventSourcedRepository struct {
	es           store.EventStore          //event store
	ca           *cache.Cache              //cache snapshot
	eventFactory event.VersionEventFactory //event factory

	//event sender
}

func (esr *EventSourcedRepository) Find(es event.EventSourced) (err error) {

	//read from cache

	var tv int64 = 0

	st := reflect.TypeOf(es).Name()
	partitionKey := GetPartitionKey(st, es.SourceId())

	teds, err := esr.es.Load(partitionKey, tv)

	if err != nil {
		return
	}

	tes := make([]event.VersionedEvent, len(teds))
	//convert to event
	for _, ted := range teds {
		var tvef event.VersionEventFactory
		te := tvef.NewVersionEvent()
		err := json.Unmarshal([]byte(ted.Payload), te)

		if err != nil {
			//todo: do extra action
			return err
		}
		tes = append(tes, te)
	}

	//load events

	for _, e := range tes {
		err = es.ApplyEvent(e)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return
	}
	return
}

func (esr *EventSourcedRepository) Save(es event.EventSourced, correlationId string) error {
	st := reflect.TypeOf(es).Name()

	partitionKey := GetPartitionKey(st, es.SourceId())

	tes := es.Events()

	eds := make([]*store.EventData, len(tes))
	for _, e := range tes {
		ed := &store.EventData{}
		ed.PartitionKey = partitionKey
		ed.SourceType = st

		//json endcode event
		payload, err := json.Marshal(e)
		if err != nil {
			return err
		}

		ed.Payload = string(payload)
		ed.SourceId = fmt.Sprintf("%d", e.SourceId())
		ed.Version = e.Version()
		eds = append(eds, ed)
	}

	//save in store
	err := esr.es.Save(partitionKey, eds)
	if err != nil {
		//todo: do extra action
		return err
	}

	//publish async

	//cache snapshot
	return nil
}

func NewRepository(es store.EventStore) *EventSourcedRepository {
	esr := &EventSourcedRepository{}
	esr.es = es

	return esr
}
