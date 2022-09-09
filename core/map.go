package core

type stdmap struct {
	storage map[string]interface{}
}

func NewMap() StorageApi {
	return &stdmap{
		storage: make(map[string]interface{}),
	}
}

func (sm *stdmap) Insert(key string, value interface{}) {
	sm.storage[key] = value
}
func (sm *stdmap) Delete(key string) {
	delete(sm.storage, key)
}
func (sm stdmap) Get(key string) (interface{}, bool) {
	val, ok := sm.storage[key]
	return val, ok
}
func (sm stdmap) Search(skip int, limit int, condition func(string, interface{}) bool) []interface{} {
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
func (sm stdmap) Length() int {
	return len(sm.storage)
}
