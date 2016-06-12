package store

//event data
type EventData struct {
	Id int64 `gorm:"primary_key"`

	SourceId      string
	Version       int64
	Payload       string
	CorrelationId string
	EventType     string
	PartitionKey  string
	SourceType    string
}

func (ed *EventData) TableName() string {
	return "t_event_data"
}

type EventStore interface {
	Load(partitionKey string, version int64) ([]*EventData, error)
	Save(partitionKey string, events []*EventData) error
}
