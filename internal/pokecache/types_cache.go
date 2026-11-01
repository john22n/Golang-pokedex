package pokecache

import (
	"sync"
	"time"
)

type Pokecache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
