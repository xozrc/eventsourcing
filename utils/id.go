package utils

import (
	"math/rand"

	"github.com/xozrc/eventsourcing/types"
)

func Guid() types.Guid {
	n := rand.Int63()
	return types.Guid(n)
}
