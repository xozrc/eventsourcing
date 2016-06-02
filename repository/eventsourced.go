package repository

import (
	"github.com/xozrc/eventsourcing/cache"
	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/store"
	"github.com/xozrc/eventsourcing/types"
	"golang.org/x/net/context"
)

// type EventSourcedRepository interface {
// 	Find(id types.Guid) (event.EventSourced, error)
// 	Save(event.EventSourced, string) error
// }

const (
	repository = "Repository"
)

func WithRepository(ctx context.Context, repo EventSourcedRepository) context.Context {
	return context.WithValue(ctx, repository, repo)
}

func RepositoryInContext(ctx context.Context) EventSourcedRepository {
	tv := ctx.Value(repository)
	tesc, _ := tv.(EventSourcedRepository)
	return tesc
}

type EventSourcedRepository struct {
	st    string            //soruce type
	es    *store.EventStore //event store
	ca    *cache.Cache      //cache snapshot
	mar   Marshaller        //event marshaller
	unmar Unmarshaller      //event unmarshaller
	 //event sender
}

func (esr *EventSourcedRepository) Find(id types.Guid, es *event.EventSourced) (err error) {

	//read from cache

	tv := 0
	partitionKey := GetPartitionKey(esr.st, id)
	teds, err := esr.es.Load(partitionKey, tv)

	if err != nil {
		return
	}

	tes := make([]event.VersionedEvent, len(teds))
	//convert to event
	for _, ted := range teds {

		te, err := esr.unmar.Unmarshal(ted)

		if err != nil {
			//todo: do extra action
			return err
		}
	}

	//load events
	err := es.LoadFrom(tes)
	if err != nil {
		return
	}

}

func (esr *EventSourcedRepository) Save(es *event.EventSourced, correlationId string) error {

	partitionKey := GetPartitionKey(esr.st, es.SourceId())

	tes := es.Events()

	eds := make([]*store.EventData, len(tes))
	for _, e := range tes {
		ed, err := esr.mar.Marshal(e)
		if err != nil {
			return err
		}
		eds = append(eds, ed)
	}

	//save in store
	err := esr.es.Save(partitionKey, eds)
	if err != nil {
		//todo: do extra action
		return
	}

	//publish async

	//cache snapshot

}

func NewRepository(es store.EventStore, sourceType string) *EventSourcedRepository {
	esr := &EventSourcedRepository{}
	esr.es = es
	esr.sourceType = sourceType
	return esr
}
