package eventsourcing

type Repository interface {
	Find(id Guid) (Aggregate, error)
	Save(Aggregate) error
}
