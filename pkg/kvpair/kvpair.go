package kvpair

import (
	"sync"
)

type Cache struct {
	*cache
}

type cache struct {
	items map[int64][]byte
	mu    sync.RWMutex
}

func newCache(m map[int64][]byte) *cache {
	c := &cache{
		items: m,
	}
	return c
}

func NewCache() *Cache {
	items := make(map[int64][]byte)
	cache := newCache(items)
	return &Cache{
		cache: cache,
	}
}

func (c *cache) Set(k int64, x []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[k] = x

}

func (c *cache) Get(k int64) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	return item, true
}

func (c *cache) GetAll() [][]byte {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var items [][]byte
	for _, item := range c.items {
		items = append(items, item)
	}
	return items
}
