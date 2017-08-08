package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShardedConcurrentMultiMap(t *testing.T) {
	concurrentMap := NewShardedConcurrentMultiMap(WithNumberOfShards(16))
	concurrentMap.Set("names", []interface{}{"Raed Shomali"})
	concurrentMap.Append("names", "Dwayne Johnson")

	ok := concurrentMap.Contains("names")
	assert.True(t, ok)

	size := concurrentMap.Size()
	assert.Equal(t, size, 1)

	value, ok := concurrentMap.Get("names")
	assert.True(t, ok)
	assert.Equal(t, value, []interface{}{"Raed Shomali", "Dwayne Johnson"})

	value, ok = concurrentMap.Get("unknown")
	assert.False(t, ok)
	assert.Nil(t, value)

	concurrentMap.Remove("names")
	value, ok = concurrentMap.Get("names")
	assert.False(t, ok)
	assert.Nil(t, value)

	concurrentMap.Append("names", "Raed Shomali")
	concurrentMap.Clear()

	value, ok = concurrentMap.Get("names")
	assert.False(t, ok)
	assert.Nil(t, value)

	size = concurrentMap.Size()
	assert.Equal(t, size, 0)
}

func TestShardedConcurrentMultiMap_InvalidShards(t *testing.T) {
	concurrentMap := NewShardedConcurrentMultiMap(WithNumberOfShards(0))
	assert.Equal(t, concurrentMap.shards, uint32(16))
}
