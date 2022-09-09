package gostore

import (
	"fmt"
	"sync/atomic"

	"github.com/Rustixir/gostore/core"
	m "github.com/Rustixir/gostore/mode"
)

var UniqueId uint32 = 0

func NextId() uint32 {
	return atomic.AddUint32(&UniqueId, 1)
}

type Storage struct {
	api     core.StorageApi
	walChan chan WalRequest
}

func NewStorage(name string, path string, bufSize uint32, mode m.Mode) Storage {
	if bufSize < 1 {
		bufSize = 10
	}
	if path == "" {
		uniqueid := NextId()
		if name == "" {
			path = fmt.Sprintf("./storage/GOS_%d", uniqueid)
		} else {
			path = fmt.Sprintf("./storage/%s_%d", name, uniqueid)
		}

	}

	walChan, err := NewWorker(path, bufSize)
	if err != nil {
		panic(err)
	}

	st := Storage{
		api:     storageBuilder(mode),
		walChan: walChan,
	}

	st.Loader()
	return st
}

func (s *Storage) Insert(key string, value interface{}) error {
	s.api.Insert(key, value)
	marshaledTuple, err := NewTuple(Insert, key, value).Marshal()
	if err != nil {
		return err
	}
	s.walChan <- marshaledTuple
	return nil
}
func (s *Storage) Delete(key string) error {
	s.api.Delete(key)
	marshaledTuple, err := NewTuple(Delete, key, 0).Marshal()
	if err != nil {
		return err
	}
	s.walChan <- marshaledTuple
	return nil
}
func (s Storage) Get(key string) (interface{}, bool) {
	val, ok := s.api.Get(key)
	return val, ok
}
func (s Storage) Search(skip int, limit int, condition func(string, interface{}) bool) []interface{} {
	collect := s.api.Search(skip, limit, condition)
	return collect
}
func (s Storage) Len() int {
	return s.api.Length()
}
func (s Storage) Shutdown() {
	req := NewShutdown()
	s.walChan <- req
	<-req.Chann
}

func (s Storage) Loader() {
	iter := NewIter(100)
	s.walChan <- iter

	counter := 0

	for {
		tuple, ok := iter.Next()

		if !ok {
			s.walChan <- SetCounter{Counter: uint64(counter)}
			return
		}

		counter++

		// Call direct api method to avoid rewriting to disk
		switch tuple.Opt {
		case Insert:
			s.api.Insert(tuple.Key, tuple.Value)
		case Delete:
			s.api.Delete(tuple.Key)
		}
	}
}

func storageBuilder(mode m.Mode) core.StorageApi {
	switch mode {
	case m.WriteConcurrency, m.ReadWriteConcurrency:
		return core.NewConcurrentMap()
	case m.ReadConcurrency:
		return core.NewSyncMap()
	case m.LockLess:
		return core.NewMap()
	default:
		return core.NewRwLockMap()
	}
}
