package mode

type Mode uint8

const (

	// std map behind a RwMutex
	RwLock Mode = iota + 1

	// std map without any locking
	LockLess

	// concurrent shared map
	WriteConcurrency

	// std sync.Map
	ReadConcurrency

	// concurrent shared map
	ReadWriteConcurrency
)
