/*
Passing channel to functions
Returning channels from function
*/

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

const (
	chCap = 10
)

func main() {
	ch := make(chan int, chCap)
	// if ch == nil(
	// 	log.Fatal("Failed to allocate memory for channel")
	// )

	//passing channels to function // shallow copy same as slice or map
	// producer(ch)
	// consumer(ch)

	d := generator(ch)
	d = counter(d)
	// consumer(d) // while reading receiver on channels makes len of channel = 0
	fmt.Printf("After adder fun...\n")
	d = adder(d, 5)
	consumer(d)
}

func adder(d chan int, add int) (out chan int) {
	out = make(chan int, len(d))
	for v := range d {
		out <- v + add
	}
	close(out)
	return
}

func counter(d chan int) (out chan int) {
	out = make(chan int, len(d))
	count := 0
	for v := range d {
		out <- v
		count++
	}
	close(out)
	fmt.Printf("Counted %v elements\n", count)
	return
}

func generator(ch chan int) (out chan int) {
	fmt.Printf("Generator of random ints...\n")
	n := r.Int()%cap(ch) + 1
	out = make(chan int, n)
	for i := 0; i < n; i++ {
		out <- r.Int()%200 + 1
	}
	close(out)
	return
}

func producer(ch chan int) {
	c := r.Int()%cap(ch) + 1
	for i := 0; i < c; i++ {
		ch <- r.Int()%200 + 1
	}
	close(ch)
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Printf("Consumer got: %v\n", v)
	}
}
