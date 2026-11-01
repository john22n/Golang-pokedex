package pokecache

import (
	"github.com/google/go-cmp/cmp"
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
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("exptected not to find key")
		return
	}
}
