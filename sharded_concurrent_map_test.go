package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShardedConcurrentMap(t *testing.T) {
	concurrentMap := NewShardedConcurrentMap(WithNumberOfShards(16))
	concurrentMap.Set("name", "Raed Shomali")

	ok := concurrentMap.Contains("name")
	assert.True(t, ok)

	size := concurrentMap.Size()
	assert.Equal(t, size, 1)

	value, ok := concurrentMap.Get("name")
	assert.True(t, ok)
	assert.Equal(t, value, "Raed Shomali")

	value, ok = concurrentMap.Get("unknown")
	assert.False(t, ok)
	assert.Nil(t, value)

	concurrentMap.Remove("name")
	value, ok = concurrentMap.Get("name")
	assert.False(t, ok)
	assert.Nil(t, value)

	concurrentMap.Set("name", "Raed Shomali")
	concurrentMap.Clear()

	value, ok = concurrentMap.Get("name")
	assert.False(t, ok)
	assert.Nil(t, value)

	size = concurrentMap.Size()
	assert.Equal(t, size, 0)
}

func TestShardedConcurrentMap_InvalidShards(t *testing.T) {
	concurrentMap := NewShardedConcurrentMap(WithNumberOfShards(0))
	assert.Equal(t, concurrentMap.shards, uint32(16))
}
