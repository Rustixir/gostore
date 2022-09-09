package gostore

import (
	"fmt"
	"testing"

	"github.com/Rustixir/gostore/mode"
)

var St Storage

func algorithm(st *Storage, t *testing.T) {

	key := "DanyalMh"
	key2 := key + "5"
	key3 := key + "3"

	value := "Senior backend developer"

	err := st.Insert(key, value)
	if err != nil {
		fmt.Println("====>", err)
	}

	err = st.Insert(key2, value)
	if err != nil {
		fmt.Println("====>", err)
	}

	err = st.Insert(key3, value)
	if err != nil {
		fmt.Println("====>", err)
	}

	_, ok := st.Get(key)
	if !ok {
		t.Errorf("key: (%s)\ninserted but not found !!!", key)
	}

	keys := []string{key2, key3}
	list := st.Search(1, 2, func(s string, i interface{}) bool {
		for _, key := range keys {
			if s == key {
				return true
			}
		}

		return false
	})

	if len(list) != 2 {
		t.Error("must get 1\nbut is 2!!!", len(list))
	}

	if list[0] != value {
		t.Errorf("expected: %v \nbut got: %v!!!", value, list)
	}

	st.Delete(key)
	_, ok = st.Get(key)
	if ok {
		t.Errorf("key: (%s)\ndeleted but not removed !!!", key)
	}
}

func TestStorageModeDefault(t *testing.T) {
	st := NewStorage("Default", "", 0, 0)
	defer st.Shutdown()
	algorithm(&st, t)

}
func TestStorageModeRwLock(t *testing.T) {
	st := NewStorage("RwLock", "", 0, mode.RwLock)
	defer st.Shutdown()
	algorithm(&st, t)

}
func TestStorageModeWriteConcurrency(t *testing.T) {
	st := NewStorage("WriteConcurrency", "", 0, mode.WriteConcurrency)
	defer st.Shutdown()
	algorithm(&st, t)

}
func TestStorageModeReadConcurrency(t *testing.T) {
	st := NewStorage("ReadConcurrency", "", 0, mode.ReadConcurrency)
	defer st.Shutdown()
	algorithm(&st, t)

}
func TestStorageModeReadWriteConcurrency(t *testing.T) {
	st := NewStorage("RwLock", "", 0, mode.ReadWriteConcurrency)
	defer st.Shutdown()
	algorithm(&st, t)

}
