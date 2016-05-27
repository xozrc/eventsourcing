package event

import "golang.org/x/net/context"

const (
	eventBus = "EventBus"
)

type EventBus interface {
	PublishEvent(e Event) error
}

func WithEventBus(ctx context.Context, eb EventBus) context.Context {
	return context.WithValue(ctx, eventBus, eb)
}

func EventBusInContext(ctx context.Context) EventBus {

	teb := ctx.Value(eventBus)
	eb, _ := teb.(EventBus)
	return eb
}
