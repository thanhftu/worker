package worker

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorker(t *testing.T) {
	var index int64 = 7
	RedisSource := "localhost:6379"
	val, err := WorkerRedisFib(index, RedisSource)
	require.NoError(t, err)
	require.Equal(t, int64(13), val)

}
