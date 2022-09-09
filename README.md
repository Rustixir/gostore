


**Gostore is a key/value storage with some different modes to**
**get better performance for that usecase, also persist data to disk to avoid loss any data**



The gostore provides the following:

* **InMemory** - whole data stored in-memory 
  bute persistent data to disk and
  after restart, **load whole data to memory**

* **Mode** - gostore have some different mode (**WriteConcurrency**, (**ReadConcurrency(**, **ReadWriteConcurrency**, **RwLock**, **Lockless**) to for choice correct hashmap data structure 
for better performance 


```
st1 := NewStorage("name_1", "./path", 0, mode.RwLock)
defer st1.Shutdown()


// Different Mode

// st := NewStorage("WriteConcurrency", "", 0, mode.WriteConcurrency)
// defer st.Shutdown()

// st := NewStorage("RwLock", "", 0, mode.ReadWriteConcurrency)
// defer st.Shutdown()


    // Insert
    err = st1.Insert(key3, value)
	if err != nil {
		// ...
	}


    // Get
	_, ok := st1.Get(key)
	if !ok {
		// ...
	}


    // Search 
	keys := []string{key2, key3}
	list := st1.Search(1, 2, func(s string, i interface{}) bool {
		for _, key := range keys {
			if s == key {
				return true
			}
		}

		return false
	})


    // Delete
	st1.Delete(key)


```