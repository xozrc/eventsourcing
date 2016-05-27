package command

import "golang.org/x/net/context"

type CommandType int

type Command interface {
	CommandType() CommandType
}

type CommandHandler interface {
	ProcessCommand(ctx context.Context, cmd Command) error
}
