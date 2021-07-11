// Accessing critical section code in go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	routineCount = 1000
)

var (
	count        uint16
	s            = rand.NewSource(time.Now().UnixNano())
	r            = rand.New(s)
	wg           sync.WaitGroup
	counterMutex sync.Mutex
	m            map[int]int
	sm           sync.Map
)

func main() {
	start := time.Now()
	m = make(map[int]int)

	for i := 0; i < routineCount; i++ {
		getRoutineCount(i)
	}

	wg.Wait()
	fmt.Printf("Total number of routine executed : %v, map : %v\n", count, m)
	value, _ := sm.Load(1)
	fmt.Printf("sm value %v, single 1 value : %v, Type : %T\n", sm, value, sm)
	fmt.Printf("Main routine took : %v\n", time.Now().Sub(start))
}

func getRoutineCount(id int) {
	// this function is called concurrently to update 'counter'
	wg.Add(1) // // need to do this before 'go' keyword
	go func() {
		// prevent concurrent update or sync critical code section
		counterMutex.Lock()
		count++           // 1. Reading , 2. writing
		m[id] = m[id] + 1 /* if we have not provided mutex lock then for editing
		//map will cause below error:
		//fatal error: concurrent map writes
		//fatal error: concurrent map read and map write */

		counterMutex.Unlock()
		sm.Store(id, id+1) // sync.Map ==> no need to add locking
		/*Map is like a Go map[interface{}]interface{} but is safe for concurrent use by
		multiple goroutines without additional locking or coordination.
		The Map type is optimized for two common use cases:
		(1) when the entry for a given key is only ever written once but read many times, as in caches that only grow, or
		(2) when multiple goroutines read, write, and overwrite entries for disjoint sets of keys.
		*/
		wg.Done()
	}()
}
