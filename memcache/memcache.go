package memcache

import (
	"errors"
	"sync"
	"time"
)

type memoryCache struct {
	ttls map[string]time.Time

	kv map[string]string
	mu sync.RWMutex
}

func newMemoryCache() *memoryCache {
	return &memoryCache{
		ttls: make(map[string]time.Time),
		kv:   make(map[string]string),
		mu:   sync.RWMutex{},
	}
}

func (m *memoryCache) Set(key, value string, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.kv[key] = value
	m.ttls[key] = time.Now().UTC().Add(ttl)

	return nil
}

func (m *memoryCache) Get(key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.kv[key]
	if !ok {
		return "", errors.New("key doesn't exist")
	}
	if m.ttls[key].After(time.Now().UTC()) {
		delete(m.kv, key)
		delete(m.ttls, key)
		return "", errors.New("key has expired")
	}

	return value, nil
}

func (m *memoryCache) Del(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.kv[key]
	if !ok {
		return errors.New("key doesn't exist")
	}
	delete(m.kv, key)
	delete(m.ttls, key)

	return nil
}
