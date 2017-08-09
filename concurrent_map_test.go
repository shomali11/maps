package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	concurrentMap := NewConcurrentMap()
	concurrentMap.Set("name", "Raed Shomali")

	ok := concurrentMap.ContainsKey("name")
	assert.True(t, ok)

	ok = concurrentMap.ContainsEntry("name", "Raed Shomali")
	assert.True(t, ok)

	size := concurrentMap.Size()
	assert.Equal(t, size, 1)

	isEmpty := concurrentMap.IsEmpty()
	assert.False(t, isEmpty)

	keys := concurrentMap.Keys()
	assert.Equal(t, keys, []string{"name"})

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

	isEmpty = concurrentMap.IsEmpty()
	assert.True(t, isEmpty)

	keys = concurrentMap.Keys()
	assert.Equal(t, keys, []string{})
}
