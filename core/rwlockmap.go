package core

import "sync"

type rwlockmap struct {
	sync.RWMutex
	storage map[string]interface{}
}

func NewRwLockMap() StorageApi {
	return &rwlockmap{
		storage: make(map[string]interface{}),
	}
}

func (sm *rwlockmap) Insert(key string, value interface{}) {
	sm.Lock()
	defer sm.Unlock()

	sm.storage[key] = value
}
func (sm *rwlockmap) Delete(key string) {
	sm.Lock()
	defer sm.Unlock()

	delete(sm.storage, key)
}
func (sm *rwlockmap) Get(key string) (interface{}, bool) {
	sm.RWMutex.RLock()
	defer sm.RUnlock()

	val, ok := sm.storage[key]
	return val, ok
}
func (sm *rwlockmap) Search(skip int, limit int, condition func(string, interface{}) bool) []interface{} {
	sm.RLock()
	defer sm.RUnlock()

	counter := 0
	collect := make([]interface{}, limit)
	for k, v := range sm.storage {
		if skip > 0 {
			skip--
			continue
		}

		if condition(k, v) {
			collect[counter] = v
			counter++
			if counter == limit {
				break
			}
		}
	}

	return collect
}
func (sm *rwlockmap) Length() int {
	return len(sm.storage)
}
