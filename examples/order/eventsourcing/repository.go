package eventsourcing


import (
	eventsourcingevent "github.com/xozrc/eventsourcing/event"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
		"github.com/xozrc/eventsourcing/repository"
)

type Repository struct {
}

func (r *Repository) Find(eventsourcingtypes.Guid) (*eventsourcingevent.EventSourced, error) {
	return nil, nil
}

func (r *Repository) Save(es *eventsourcingevent.EventSourced) error {
	return nil
}

func RepositoryInContext(ctx context.Context) *Repository{
	tr :=repository.RepositoryInContext(ctx)
	if tr ==nil{
		return nil
	}
	tt,_:=tr.(*Repository)
	return tt
}