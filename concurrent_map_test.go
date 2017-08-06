package cmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	concurrentMap := NewConcurrentMap()
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
