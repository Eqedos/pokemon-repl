package cache

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	c := New(5 * time.Minute)

	// Test adding and retrieving data
	key := "test-key"
	data := []byte("test-data")

	c.Add(key, data)

	got, ok := c.Get(key)
	if !ok {
		t.Error("expected to find cached data, but got nothing")
	}

	if string(got) != string(data) {
		t.Errorf("expected %q, got %q", string(data), string(got))
	}
}

func TestCacheGetMiss(t *testing.T) {
	c := New(5 * time.Minute)

	// Test getting non-existent key
	got, ok := c.Get("non-existent")
	if ok {
		t.Errorf("expected cache miss, but got data: %q", string(got))
	}
}

func TestCacheOverwrite(t *testing.T) {
	c := New(5 * time.Minute)

	key := "test-key"
	c.Add(key, []byte("first"))
	c.Add(key, []byte("second"))

	got, ok := c.Get(key)
	if !ok {
		t.Error("expected to find cached data")
	}

	if string(got) != "second" {
		t.Errorf("expected %q, got %q", "second", string(got))
	}
}
