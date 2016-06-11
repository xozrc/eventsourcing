package command

import (
	"github.com/xozrc/eventsourcing/command"
	"golang.org/x/net/context"
)

type OrderCreateCommand struct {
}

func OrderCreate(ctx context.Context, cmd command.Command) (err error) {
	return nil
}
