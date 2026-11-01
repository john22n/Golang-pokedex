package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Pokecache {
	var cache Pokecache
	cache.cache = make(map[string]cacheEntry)
	go cache.reapLoop(interval)

	return &cache
}

func (p *Pokecache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		p.mu.Lock()
		for key, entry := range p.cache {
			if time.Since(entry.createdAt) > interval {
				delete(p.cache, key)
			}
		}
		p.mu.Unlock()
	}
}

func (p *Pokecache) Add(key string, val []byte) {
	p.mu.Lock()
	p.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	p.mu.Unlock()
}

func (p *Pokecache) Get(key string) ([]byte, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entry, ok := p.cache[key]

	if !ok {
		return nil, false
	}
	return entry.val, true
}
