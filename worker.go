package gostore

import (
	w "github.com/tidwall/wal"
)

func NewWorker(path string, bufSize uint32) (chan WalRequest, error) {

	log, err := w.Open(path, nil)
	if err != nil {
		return nil, err
	}

	chann := make(chan WalRequest, bufSize)
	go start(log, chann)

	return chann, nil
}

func start(log *w.Log, chann chan WalRequest) {
	go func() {
		var counter uint64

		for req := range chann {
			switch request := req.(type) {
			case SetCounter:
				counter = request.Counter + 1
			case MarshaledTuple:
				log.Write(counter, request.Bytes)
				counter += 1

			case Shutdown:
				log.Close()
				request.Chann <- struct{}{}

			case Iter:
				var reader uint64 = 1
				for {
					bytes, err := log.Read(reader)
					if err != nil {
						close(request.Chann)
						break
					}

					reader++
					tuple, _ := TupleFromBytes(bytes)
					request.Chann <- tuple
				}
			}
		}
	}()
}
