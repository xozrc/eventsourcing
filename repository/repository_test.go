package repository_test

import (
	"errors"
	"math/rand"
	"reflect"
	"sync"
	"testing"
)
import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xozrc/eventsourcing/event"
	"github.com/xozrc/eventsourcing/event/mock_event"
	. "github.com/xozrc/eventsourcing/repository"
	"github.com/xozrc/eventsourcing/store"
	"github.com/xozrc/eventsourcing/store/mock_store"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

var (
	findCorrectId     = eventsourcingtypes.Guid(1)
	findNoExistId     = eventsourcingtypes.Guid(2)
	findVersionMissId = eventsourcingtypes.Guid(3)
	findStoreErrorId  = eventsourcingtypes.Guid(4)
)

var (
	once sync.Once
)

var (
	minEvents = 3
	maxEvents = 10
	versions  []event.VersionedEvent
)

var (
	storeErr = errors.New("load store error")
)

func setup() {
	once.Do(func() {
		n := rand.Intn(maxEvents)
		if n < minEvents {
			n = minEvents
		}
		versions = make([]event.VersionedEvent, 0, n)

		for i := 0; i < n; i++ {
			tv := event.NewVersionEvent(eventsourcingtypes.Guid(int64(i+1)), int64(i+1))
			versions = append(versions, tv)
		}
	})
}

func TestLoadError(t *testing.T) {
	testLoad(t, findStoreErrorId)
}

func TestLoadCorrect(t *testing.T) {
	testLoad(t, findCorrectId)
}

func testLoad(t *testing.T, testId eventsourcingtypes.Guid) {
	setup()
	//mock event sourced
	ectrl := gomock.NewController(t)
	tes1 := mock_event.NewMockEventSourced(ectrl)
	defer ectrl.Finish()

	sourceType := reflect.TypeOf(tes1).Name()
	partitionKey := GetPartitionKey(sourceType, testId)

	tevents, loadErr := eventsForSourceId(testId)

	tes1.EXPECT().SourceId().Return(testId).AnyTimes()

	eds := make([]*store.EventEntity, 0)

	for _, tevent := range tevents {
		tes1.EXPECT().ApplyEvent(tevent).Return(nil)
		tEventEntity, _ := ToData(sourceType, partitionKey, tevent)
		eds = append(eds, tEventEntity)
	}

	var finalVersion int64 = 0
	if len(tevents) >= 1 {
		finalVersion = tevents[len(tevents)-1].Version()
		tes1.EXPECT().Version().Return(finalVersion)
	}

	//mock event store
	ctrl := gomock.NewController(t)
	tes := mock_store.NewMockEventStore(ctrl)
	defer ctrl.Finish()

	esr := NewRepository(tes, nil)

	tes.EXPECT().Load(partitionKey, int64(0)).Return(eds, loadErr)

	err := esr.Find(tes1)
	if loadErr != nil {
		if !assert.EqualError(t, err, loadErr.Error()) {
			return
		}
		return
	}
	if !assert.NoError(t, err, "store error") {
		return
	}
	if !assert.Equal(t, tes1.Version(), finalVersion, "version no equal") {
		return
	}

}

func eventsForSourceId(id eventsourcingtypes.Guid) ([]event.VersionedEvent, error) {

	if id == findCorrectId {
		return versions, nil
	}

	if id == findStoreErrorId {
		return nil, storeErr
	}

	return nil, nil
}
