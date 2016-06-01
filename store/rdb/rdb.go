package rdb

import "github.com/xozrc/eventsourcing/repository"

type RdbEventStore struct {
}

func (res *RdbEventStore) Load(partitionKey string, version int64) []*repository.EventData {
	return
}

func (res *RdbEventStore) Save(partitionKey string, events []*EventData) error {
	return
}
