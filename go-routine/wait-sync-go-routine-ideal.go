// waiting and sync of go routine using ideal way

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
	fmt.Println("--Waiting and sync on go routines--")

	start := time.Now()

	wg.Add(3) // wait for 3 go routine to complete
	go Producer(1)
	go Producer(2)
	go Producer(3)

	wg.Wait()

	wg.Add(1)
	go Producer(4)
	wg.Wait()
	elapse := time.Now().Sub(start)
	fmt.Printf("Ideal wait took %v\n", elapse)
}

func Producer(no int) {
	n := r.Int()%1000 + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer # %v ran for %v\n", no, d)
	wg.Done()
}
