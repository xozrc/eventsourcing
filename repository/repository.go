package repository

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

import (
	"github.com/xozrc/eventsourcing/cache"
	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/store"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

type Repository interface {
	Find(id eventsourcingtypes.Guid, es event.EventSourced) (err error)
	Save(es event.EventSourced, correlationId string) error
}

type EventSourcedRepository struct {
	es store.EventStore //event store
	ca cache.Cache      //cache snapshot
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

	tes := make([]event.VersionedEvent, 0, len(teds))
	//convert to event
	for _, ted := range teds {
		factory := event.GetVersionEventFactory(ted.EventType)
		if factory == nil {
			return errors.New("no found factory")
		}
		sourceId, err := strconv.ParseInt(ted.SourceId, 10, 64)
		if err != nil {
			return err
		}

		te := factory.NewVersionEvent(eventsourcingtypes.Guid(sourceId), ted.Version)
		err = json.Unmarshal([]byte(ted.Payload), te)

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

	eds := make([]*store.EventEntity, len(tes))
	for _, e := range tes {
		ed, err := ToData(st, partitionKey, e)
		if err != nil {
			return err
		}
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

func NewRepository(es store.EventStore, ca cache.Cache) *EventSourcedRepository {
	esr := &EventSourcedRepository{}
	esr.es = es
	esr.ca = ca
	return esr
}
