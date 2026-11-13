package pokecache

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := map[string]struct {
		key string
		val []byte
	}{
		"test root url": {
			key: "https://example.com",
			val: []byte("testdata"),
		},
		"test route pat": {
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, tc.val)
			val, ok := cache.Get(tc.key)

			if !ok {
				t.Errorf("expected to find key")
				return
			}
			t.Logf("diff: %s", cmp.Diff(val, tc.val))
			if cmp.Diff(val, tc.val) != "" {
				t.Errorf("expected to find value")
				return
			}
		})
	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if _, ok := cache.Get("https://example.com"); !ok {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatal("expected entry to expire within one second")
}

func TestConcurrentAddGet(t *testing.T) {
	cache := NewCache(time.Hour)
	var workers sync.WaitGroup
	for worker := 0; worker < 16; worker++ {
		workers.Add(1)
		go func(worker int) {
			defer workers.Done()
			key := fmt.Sprintf("key-%d", worker)
			for version := 0; version < 50; version++ {
				value := []byte(fmt.Sprintf("value-%d-%d", worker, version))
				cache.Add(key, value)
				got, ok := cache.Get(key)
				if !ok || string(got) != string(value) {
					t.Errorf("cache lost value for %s", key)
					return
				}
			}
		}(worker)
	}
	workers.Wait()
	if _, ok := cache.Get("missing"); ok {
		t.Fatal("missing key must not be present")
	}
}
