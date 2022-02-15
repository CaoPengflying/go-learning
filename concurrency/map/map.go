// @Description _map
// @Author caopengfei 2021/2/24 11:26
package main

import (
	"sync"
)

func main() {
	m := sync.Map{}
}

type ConcurrentMap []*ConcurrentMapShard

const SHARD_COUNT = 32

type ConcurrentMapShard struct {
	sync.RWMutex
	items map[string]interface{}
}

func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShard{items: make(map[string]interface{})}
	}
	return m
}

func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShard {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}
func (m ConcurrentMap) Set(key string, value interface{}) {
	shard := m.GetShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.items[key] = value
}
func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	shard := m.GetShard(key)
	shard.RLock()
	defer shard.RUnlock()
	value, exist := shard.items[key]
	return value, exist
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

type RWMap struct {
	sync.RWMutex
	m map[int]int
}

func (r *RWMap) Get(key int) (int, bool) {
	r.RLock()
	defer r.RUnlock()
	value, exist := r.m[key]
	return value, exist
}

func (r *RWMap) Set(key, value int) {
	r.Lock()
	defer r.Unlock()
	r.m[key] = value
}
