package generics_test

import (
	"testing"

	"github.com/quentinlesceller/generics"
	"github.com/stretchr/testify/assert"
)

func TestStoreLoad(t *testing.T) {
	var m generics.SyncMap[string, int]
	key := "test"
	value := 1
	m.Store(key, value)
	actualValue, ok := m.Load(key)
	assert.Equal(t, value, actualValue)
	assert.True(t, ok)
}

func TestLoadOrStore(t *testing.T) {
	var m generics.SyncMap[string, int]
	key := "test"
	value := 1
	m.Store(key, value)
	m.LoadOrStore(key, value)
	actualValue, ok := m.Load(key)
	assert.Equal(t, value, actualValue)
	assert.True(t, ok)
}

func TestLoadAndDelete(t *testing.T) {
	var m generics.SyncMap[string, int]
	key := "test"
	value := 1
	m.Store(key, value)
	actualValue, ok := m.LoadAndDelete(key)
	assert.Equal(t, value, actualValue)
	assert.True(t, ok)
	actualValue, ok = m.Load(key)
	assert.Equal(t, 0, actualValue)
	assert.False(t, ok)
}

func TestDelete(t *testing.T) {
	var m generics.SyncMap[string, int]
	key := "test"
	value := 1
	m.Store(key, value)
	m.Delete(key)
	actualValue, ok := m.Load(key)
	assert.Equal(t, 0, actualValue)
	assert.False(t, ok)
}

func TestRange(t *testing.T) {
	var m generics.SyncMap[string, int]
	keys := []string{"hello", "world"}
	values := []int{1, 2}
	m.Store(keys[0], values[0])
	m.Store(keys[1], values[1])
	// range does not guarantee ordering
	m.Range(func(k string, v int) bool {
		index := 0
		if k == keys[1] {
			index = 1
		}
		assert.Equal(t, keys[index], k)
		assert.Equal(t, values[index], v)
		index++
		return true
	})
}
