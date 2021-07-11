// waiting and synchronization of go routine
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	s = rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
)

func main() {
	fmt.Println("--Waiting and sync on go routines--")
	start := time.Now()

	go Producer(1)
	go Producer(2)

	time.Sleep(1 * time.Millisecond) // waiting arb time to complete execution of go routine
	elapse := time.Now().Sub(start)
	fmt.Printf("Non-ideal wait took %v\n", elapse)
}

func Producer(no int) {
	n := r.Int()%1000 + 1 // generate random numbers from 1 to 1000
	d := time.Duration(n)
	time.Sleep(d)
	fmt.Printf("Producer # %v ran for %v\n", no, d)
}
