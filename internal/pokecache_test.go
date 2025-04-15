package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("example"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expexted to find key %s in cache", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected %s, got %s", string(c.val), string(val))
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	fmt.Println("Testing reap loop")
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("example"))
	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key %s in cache", "https://example.com")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key %s in cache", "https://example.com")
		return
	}
}
