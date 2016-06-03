#! /bin/bash -e

mockgen github.com/xozrc/eventsourcing/event EventSourced \
  > event/mock_event/mock_event.go
  
mockgen github.com/xozrc/eventsourcing/store EventStore \
  > store/mock_store/mock_store.go
gofmt -w store/mock_store/mock_store.go event/mock_event/mock_event.go

echo >&2 "OK"
