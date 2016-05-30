package repository

import (
	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/types"
	"golang.org/x/net/context"
)

type EventSourcedRepository interface {
	Find(id types.Guid) (event.EventSourced, error)
	Save(event.EventSourced, string) error
}

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
