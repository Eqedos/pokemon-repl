// Package cache provides a thread-safe key-value store with TTL-based expiration.
package cache

import (
	"sync"
	"time"
)

// Cache provides a thread-safe key-value store with automatic TTL-based expiration.
// It uses a read-write mutex to allow concurrent reads while ensuring safe writes.
type Cache struct {
	entries map[string]entry
	mu      *sync.RWMutex
	ttl     time.Duration
}

// entry represents a single cached item with its creation timestamp.
type entry struct {
	createdAt time.Time
	data      []byte
}

// New creates a new Cache instance with the specified TTL duration.
// A background goroutine is started to automatically remove expired entries.
func New(ttl time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]entry),
		mu:      &sync.RWMutex{},
		ttl:     ttl,
	}
	go c.reapLoop()
	return c
}

// Add stores a value in the cache with the given key.
// If the key already exists, its value is overwritten.
func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = entry{
		createdAt: time.Now(),
		data:      data,
	}
}

// Get retrieves a value from the cache by key.
// Returns the value and true if found, or nil and false if not present.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return e.data, true
}

// reapLoop runs in a background goroutine to periodically remove expired entries.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, e := range c.entries {
			if time.Since(e.createdAt) > c.ttl {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
