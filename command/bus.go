package command

import "golang.org/x/net/context"

const (
	commandBus = "CommandBus"
)

func WithCommandBus(ctx context.Context, cb CommandBus) context.Context {
	return context.WithValue(ctx, commandBus, cb)
}

func CommandBusInContext(ctx context.Context) CommandBus {

	tcb := ctx.Value(commandBus)
	tb, _ := tcb.(CommandBus)
	return tb
}

type CommandBus interface {
	SendCommand(cmd Command) error
}
