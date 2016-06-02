package store

type EventData struct {
	Id            int64 `gorm:"primary_key"`
	RowKey        string
	PartitionKey  string
	SourceId      string
	Version       int64
	SourceType    string
	Payload       string
	CorrelationId string
}

func (ed *EventData) TableName() string {
	return "t_event_data"
}

type EventStore interface {
	Load(partitionKey string, version int64) ([]*EventData, error)
	Save(partitionKey string, events []*EventData) error
}
