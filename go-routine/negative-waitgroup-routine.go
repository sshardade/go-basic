package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s  = rand.NewSource(time.Now().UnixNano())
	r  = rand.New(s)
	wg sync.WaitGroup
)

func main() {
	start := time.Now()

	// Add adds delta, which may be negative, to the WaitGroup counter.
	// If the counter becomes zero, all goroutines blocked on Wait are released.
	// If the counter goes negative, Add panics.
	wg.Add(2) // if we add here negative waitgroup counter then also will get panic
	go Producer(1)
	go Producer(2)

	wg.Wait()
	elapse := time.Now().Sub(start)
	fmt.Printf("Time took by main routine : %v\n", elapse)
}

func Producer(id int) {
	start := time.Now()
	t := r.Int()%1000 + 1
	d := time.Duration(t) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer # %v - took %v\n", id, time.Now().Sub(start))
	//wg.Done()
	wg.Add(-1) /* Instead of wg.Done() we can also use wg.Add(-1) as
		internal implementation also using same add call with negative counter
	// Done decrements the WaitGroup counter by one.
	func (wg *WaitGroup) Done() {
		wg.Add(-1)
	}*/
}
