package markov

import (
	"sync"
)

// mem is the in memory file store
type mem struct {
	cache map[string]map[string]int
	*sync.Mutex
}

func NewMemStore() ModelStore {
	return &mem{
		cache: make(map[string]map[string]int),
		Mutex: &sync.Mutex{},
	}
}

func (mem *mem) Set(key string, counts map[string]int) error {
	mem.Lock()
	defer mem.Unlock()
	mem.cache[key] = counts
	return nil
}

func (mem *mem) Get(key string) (map[string]int, error) {
	mem.Lock()
	defer mem.Unlock()
	counts, ok := mem.cache[key]
	if ok {
		return counts, nil
	}
	return make(map[string]int), nil
}
