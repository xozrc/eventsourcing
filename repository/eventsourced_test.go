package repository_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xozrc/eventsourcing/event/mock_event"
	. "github.com/xozrc/eventsourcing/repository"
	"github.com/xozrc/eventsourcing/store/mock_store"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

var (
	findCorrectId     eventsourcingtypes.Guid = eventsourcingtypes.Guid(1)
	findNoExistId     eventsourcingtypes.Guid = eventsourcingtypes.Guid(2)
	findVersionMissId eventsourcingtypes.Guid = eventsourcingtypes.Guid(3)
	findStoreErrorId  eventsourcingtypes.Guid = eventsourcingtypes.Guid(4)
)

var (
	storeErr = errors.New("load store error")
)

func TestLoadError(t *testing.T) {

	//mock event store
	ctrl := gomock.NewController(t)
	tes := mock_store.NewMockEventStore(ctrl)
	defer ctrl.Finish()

	sourceType := ""
	esr := NewRepository(tes, sourceType)

	partitionKey := GetPartitionKey(sourceType, findStoreErrorId)

	tes.EXPECT().Load(partitionKey, int64(0)).Return(nil, storeErr)

	//mock event sourced
	ectrl := gomock.NewController(t)
	tes1 := mock_event.NewMockEventSourced(ectrl)
	defer ectrl.Finish()

	tes1.EXPECT().SourceId().Return(findStoreErrorId).AnyTimes()

	err := esr.Find(tes1)
	if !assert.EqualError(t, err, storeErr.Error(), "store error") {
		return
	}

}

// var (
// 	findCorrectId     eventsourcingtypes.Guid = eventsourcingtypes.Guid(1)
// 	findNoExistId     eventsourcingtypes.Guid = eventsourcingtypes.Guid(2)
// 	findVersionMissId eventsourcingtypes.Guid = eventsourcingtypes.Guid(3)
// 	findStoreError    eventsourcingtypes.Guid = eventsourcingtypes.Guid(4)
// )

// var (
// 	storeErr = errors.New("load store error")
// )

// var (
// 	sourceId = eventsourcingtypes.Guid(int64(1))
// )

// var (
// 	firstVersion  = event.NewVersionEvent(sourceId, 1)
// 	secondVersion = event.NewVersionEvent(sourceId, 2)
// 	thirdVersion  = event.NewVersionEvent(sourceId, 3)
// )

// var (
// 	firstVersionEventData  = &store.EventData{}
// 	secondVersionEventData = &store.EventData{}
// 	thirdVersionEventData  = &store.EventData{}
// )

// var (
// 	testDataMap map[eventsourcingtypes.Guid][]*store.EventData
// )

// func setup() {
// 	testDataMap = make(map[eventsourcingtypes.Guid][]*store.EventData, 0)
// 	testDataMap[findCorrectId] = []*store.EventData{firstVersionEventData, secondVersionEventData, thirdVersionEventData}
// 	testDataMap[findNoExistId] = []*store.EventData{}
// 	testDataMap[findVersionMissId] = []*store.EventData{firstVersionEventData, thirdVersionEventData}
// }

// func TestFind(t *testing.T) {

// }

// func TestFindNoExist(t *testing.T) {

// }

// func TestFindVersionMissId(t *testing.T) {

// }

// func testFind(t *testing.T, es event.EventSourced) (err error) {

// 	ctrl := gomock.NewController(t)
// 	tes := mock_store.NewMockEventStore(ctrl)
// 	defer ctrl.Finish()

// 	sourceType := ""
// 	esr := NewRepository(tes, sourceType)

// 	paritionKey := GetPartitionKey(sourceType, es.SourceId())
// 	version := es.Version()

// 	//store event datas
// 	storeEventDatas, storeErr := eventDatas(es.SourceId(), paritionKey, version)

// 	tes.EXPECT().Load(paritionKey, version).Return(storeEventDatas, storeErr)

// 	err = esr.Find(es.SourceId(), es)
// 	if err != nil {
// 		return
// 	}
// 	return

// }

// func eventDatas(sourceId eventsourcingtypes.Guid, partitionKey string, version int64) ([]*store.EventData, error) {
// 	tmp, ok := testDataMap[sourceId]
// 	if !ok {
// 		return nil, storeErr
// 	}
// 	return tmp, nil
// }

// func TestSave(t *testing.T) {

// }

// //mock store and mock eventsourced

// func TestFind2(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	tes := mock_store.NewMockEventStore(ctrl)
// 	defer ctrl.Finish()

// }
