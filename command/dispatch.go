package command

import (
	"errors"

	"golang.org/x/net/context"
)

var (
	CommandHandlerNoFound = errors.New("command handler no found")
)

const (
	initCommandHandlerSize = 10
)

type CommandDispatcher interface {
	DispatchCommand(ctx context.Context, c Command) error
	Register(ct CommandType, ch CommandHandler)
	Deregister(ct CommandType)
}

type commandDispatcher struct {
	handlersMap map[CommandType]CommandHandler
}

func (cd *commandDispatcher) DispatchCommand(ctx context.Context, c Command) error {
	h, ok := cd.handlersMap[c.CommandType()]
	if !ok {
		return CommandHandlerNoFound
	}
	return h.ProcessCommand(ctx, c)
}

func (cd *commandDispatcher) Register(ct CommandType, ch CommandHandler) {
	cd.handlersMap[ct] = ch
}

func (cd *commandDispatcher) Deregister(ct CommandType) {
	delete(cd.handlersMap, ct)
}

func NewCommandDispatcher() CommandDispatcher {
	cd := &commandDispatcher{}
	cd.handlersMap = make(map[CommandType]CommandHandler, initCommandHandlerSize)
	return cd
}
