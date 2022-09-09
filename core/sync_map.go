package core

import "sync"

type syncmap struct {
	storage sync.Map
}

func NewSyncMap() StorageApi {
	return &syncmap{
		storage: sync.Map{},
	}
}

func (sm *syncmap) Insert(key string, value interface{}) {
	sm.storage.Store(key, value)
}
func (sm *syncmap) Delete(key string) {
	sm.storage.Delete(key)
}
func (sm syncmap) Get(key string) (interface{}, bool) {
	return sm.storage.Load(key)
}
func (sm syncmap) Search(skip int, limit int, condition func(string, interface{}) bool) []interface{} {
	counter := 0
	collect := make([]interface{}, limit)

	sm.storage.Range(func(key, value any) bool {
		if skip > 0 {
			skip--
			return true
		}

		if condition(key.(string), value) {
			collect[counter] = value
			counter++
			if counter == limit {
				return false
			}
		}

		return true
	})

	return collect
}
func (sm syncmap) Length() int {
	len := 0
	sm.storage.Range(func(key, value any) bool {
		len++
		return true
	})
	return len
}
