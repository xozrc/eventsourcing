package eventsourcing

type Guid int64

type AggregateRoot interface {
	Guid() Guid
}

type Aggregate interface {
	AggregateRoot
	LoadFrom()
}
