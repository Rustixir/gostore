package gostore

import (
	"encoding/json"
)

type OperationType = uint8

const (
	Insert OperationType = iota + 1
	Delete
)

type Tuple struct {
	Opt   OperationType
	Key   string
	Value interface{}
}

func NewTuple(opt OperationType, key string, value interface{}) Tuple {
	return Tuple{
		Opt:   opt,
		Key:   key,
		Value: value,
	}
}
func TupleFromBytes(bytes []byte) (tuple Tuple, err error) {
	err = json.Unmarshal(bytes, &tuple)
	return tuple, err
}
func (t Tuple) Marshal() (MarshaledTuple, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return MarshaledTuple{}, ErrMarshalFailed
	}
	return MarshaledTuple{Bytes: bytes}, nil
}
