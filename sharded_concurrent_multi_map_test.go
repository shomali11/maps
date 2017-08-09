package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShardedConcurrentMultiMap(t *testing.T) {
	concurrentMap := NewShardedConcurrentMultiMap(WithNumberOfShards(16))
	concurrentMap.Set("names", []interface{}{"Raed Shomali"})
	concurrentMap.Append("names", "Dwayne Johnson")

	ok := concurrentMap.ContainsKey("names")
	assert.True(t, ok)

	ok = concurrentMap.ContainsEntry("name", "Raed Shomali")
	assert.True(t, ok)

	size := concurrentMap.Size()
	assert.Equal(t, size, 1)

	isEmpty := concurrentMap.IsEmpty()
	assert.False(t, isEmpty)

	keys := concurrentMap.Keys()
	assert.Equal(t, keys, []string{"names"})

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

	isEmpty = concurrentMap.IsEmpty()
	assert.True(t, isEmpty)

	keys = concurrentMap.Keys()
	assert.Equal(t, keys, []string{})
}

func TestShardedConcurrentMultiMap_InvalidShards(t *testing.T) {
	concurrentMap := NewShardedConcurrentMultiMap(WithNumberOfShards(0))
	assert.Equal(t, concurrentMap.shards, uint32(16))
}
