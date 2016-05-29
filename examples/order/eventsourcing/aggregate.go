package eventsourcing

type OrderStatus int

type Order struct {
	Id     Guid
	Status OrderStatus
}

func NewOrder() *Order {

}
