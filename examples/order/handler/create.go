package handler

import (
	"github.com/xozrc/eventsourcing/command"
	ordereventsourcing "github.com/xozrc/eventsourcing/examples/order/eventsourcing"
	"golang.org/x/net/context"
)

func OrderCreate(ctx context.Context, cmd command.Command) error {
	r := ordereventsourcing.RepositoryInContext(ctx)
	if r == nil {
		return nil
	}

	//create new order
	o := ordereventsourcing.NewOrder()
	//save order
	err := r.Save(o)

	return err
}
