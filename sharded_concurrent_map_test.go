package cmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShardedConcurrentMap(t *testing.T) {
	concurrentMap := NewShardedConcurrentMap(WithNumberOfShards(16))
	concurrentMap.Set("name", "Raed Shomali")

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
}

func TestShardedConcurrentMap_InvalidShards(t *testing.T) {
	concurrentMap := NewShardedConcurrentMap(WithNumberOfShards(0))
	assert.Equal(t, concurrentMap.numberOfShards, uint32(16))
}
