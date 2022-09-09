package core

import cp "github.com/orcaman/concurrent-map"

type concmap struct {
	storage cp.ConcurrentMap
}

func NewConcurrentMap() StorageApi {
	return &concmap{
		storage: cp.New(),
	}
}

func (sm *concmap) Insert(key string, value interface{}) {
	sm.storage.Set(key, value)
}
func (sm *concmap) Delete(key string) {
	sm.storage.Remove(key)
}
func (sm concmap) Get(key string) (interface{}, bool) {
	val, ok := sm.storage.Get(key)
	return val, ok
}
func (sm concmap) Search(skip int, limit int, condition func(string, interface{}) bool) []interface{} {
	counter := 0
	collect := make([]interface{}, limit)

	for tuple := range sm.storage.IterBuffered() {
		if skip > 0 {
			skip--
			continue
		}

		if condition(tuple.Key, tuple.Val) {
			collect[counter] = tuple.Val
			counter++
			if counter == limit {
				break
			}
		}
	}

	return collect
}
func (sm concmap) Length() int {
	return sm.storage.Count()
}
