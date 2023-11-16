package storage

import "sync"

// MemoryStore is struct which holds data in in-memory
type MemoryStore struct {
	mu   *sync.RWMutex
	data [][]byte
}

// NewMemoryStore is function to initiate memory storage implementation
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

// Produce Method Implementation for Memory Store
func (ms *MemoryStore) Produce(b []byte) (int, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.data = append(ms.data, b)
	return len(ms.data) - 1, nil
}

// Consume Methoda Implementation for Memory Store
func (ms *MemoryStore) Consume(offset int) ([]byte, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	return ms.data[offset], nil
}
