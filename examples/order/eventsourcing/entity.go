package eventsourcing

type OrderStatus int

const (
	OrderInit OrderStatus = iota
	OrderPending
	OrderConfirm
	OrderVerify
	OrderCancelled
)

//order entity
type OrderEntity struct {
	Id         int64
	Status     int
	Deleted    int
	UpdateTime int64
	CreateTime int64
}

//order object

type Order struct {
	Id         int64
	Status     OrderStatus
	Deleted    int
	UpdateTime int64
	CreateTime int64
}

func (o *Order) Cancel() error {
	o.Status = OrderCancelled
	return nil
}

func (o *Order) Pending() error {
	o.Status = OrderPending
	return nil
}

func (o *Order) Verify() error {
	o.Status = OrderVerify
	return nil
}
