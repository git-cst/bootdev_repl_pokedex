package pokecache_test

import (
	"bytes"
	"sync"
	"testing"
	"time"

	"github.com/git-cst/bootdev_pokedex/internal/pokecache"
)

func TestAddToCache(t *testing.T) {
	cases := []struct {
		name     string // Test case name
		key      string // Key to add
		value    []byte // Value to add
		expected []byte // Expected value in cache
	}{
		{
			name:     "Add single string value",
			key:      "test_key",
			value:    []byte("test_value"),
			expected: []byte("test_value"),
		},
		{
			name:     "Add empty value",
			key:      "empty_key",
			value:    []byte{},
			expected: []byte{},
		},
		{
			name:     "Add binary data",
			key:      "binary_key",
			value:    []byte{0x00, 0x01, 0x02, 0x03},
			expected: []byte{0x00, 0x01, 0x02, 0x03},
		},
		{
			name:     "Overwrite existing key",
			key:      "overwrite_key",
			value:    []byte("new_value"),
			expected: []byte("new_value"),
		},
	}

	// Run test cases
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Initialize a new cache for each test
			cache := pokecache.NewCache()

			// For the overwrite test, first add an initial value
			if tc.name == "Overwrite existing key" {
				cache.Add(tc.key, []byte("old_value"))
			}

			// Add the value to the cache
			beforeAdd := time.Now()
			cache.Add(tc.key, tc.value)
			afterAdd := time.Now()

			// Verify the entry exists
			entry, exists := cache.Entries[tc.key]
			if !exists {
				t.Errorf("Expected key %s to exist in cache, but it doesn't", tc.key)
			}

			// Verify the value is correct
			if !bytes.Equal(entry.Val, tc.expected) {
				t.Errorf("Expected value %v, got %v", tc.expected, entry.Val)
			}

			// Verify the timestamp is reasonable (between before and after the Add call)
			if entry.CreatedAt.Before(beforeAdd) || entry.CreatedAt.After(afterAdd) {
				t.Errorf("Expected CreatedAt timestamp to be between %v and %v, got %v",
					beforeAdd, afterAdd, entry.CreatedAt)
			}
		})
	}
}

func TestReapCache(t *testing.T) {
	cache := pokecache.Cache{
		Entries:  make(map[string]pokecache.CacheEntry),
		Interval: time.Duration(1 * time.Nanosecond),
		Mutex:    sync.Mutex{},
	}

	testKey := "Test"
	testVal := []byte("Value")
	cache.Add(testKey, testVal)

	_, exists := cache.Entries[testKey]
	if !exists {
		t.Errorf("Expected key %s to exist in cache, but it doesn't", testKey)
	}

	time.Sleep(2 * time.Nanosecond)

	cache.Reap()

	_, exists = cache.Entries[testKey]
	if exists {
		t.Errorf("Expected key %s to be removed after reaping, but it still exists", testKey)
	}
}
