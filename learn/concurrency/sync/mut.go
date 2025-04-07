package sync

import "sync"

type SafeMap struct {
	mu sync.RWMutex
	data map[string]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap {
		data: make(map[string]interface{}),
	}
}

func (m *SafeMap) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *SafeMap) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.data[key]
	return val, ok
}