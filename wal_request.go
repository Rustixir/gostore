package gostore

import "errors"

type WalRequest interface {
	Request()
}
type SetCounter struct {
	Counter uint64
}
type MarshaledTuple struct {
	Bytes []byte
}
type Iter struct {
	Chann chan Tuple
}
type Shutdown struct {
	Chann chan struct{}
}

func (sc SetCounter) Request() {}

func (mt MarshaledTuple) Request() {}

func (t Iter) Request() {}

func (s Shutdown) Request() {}

func NewShutdown() Shutdown {
	return Shutdown{
		Chann: make(chan struct{}),
	}
}

func NewIter(bufSize int) Iter {
	return Iter{
		Chann: make(chan Tuple, bufSize),
	}
}

func (it Iter) Next() (Tuple, bool) {
	tup, ok := <-it.Chann
	return tup, ok
}

var (
	ErrMarshalFailed   error = errors.New("marshaling failed")
	ErrUnmarshalFailed error = errors.New("unmarshaling failed")
)
