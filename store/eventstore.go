package store

type EventData struct {
	Id            int64
	SourceId      string
	Version       int64
	SourceType    string
	Payload       string
	CorrelationId string
}

type EventStore interface {
	Load(partitionKey string, version int64) ([]*EventData, error)
	Save(partitionKey string, events []*EventData) error
}
