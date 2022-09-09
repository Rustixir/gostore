package core

type StorageApi interface {
	Insert(key string, value interface{})
	Delete(key string)
	Get(key string) (interface{}, bool)
	Search(skip int, limit int, condition func(string, interface{}) bool) []interface{}
	Length() int
}
