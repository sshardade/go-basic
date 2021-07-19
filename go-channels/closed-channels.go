/*
1. Working on 'closed' channel
2. Range operator
*/

package main

import "fmt"

func main() {
	sc := make(chan string, 2)
	sc <- "suraj"
	close(sc) // closing channel
	fmt.Printf("ch : %v, Type: %T, len : %v, cap : %v\n", sc, sc, len(sc), cap(sc))

	//Sending on closed channel
	// sc <- "Hardade" //  panic: send on closed channel

	//Receive from closed channel
	ch1 := make(chan string, 2)
	ch1 <- "Suraj"
	close(ch1)
	ch1Received := <-ch1
	fmt.Printf("ch1 1st value: %v, len : %v, cap : %v\n", ch1Received, len(ch1), cap(ch1)) // "Suraj", len : 0, cap: 2
	ch1Received = <-ch1
	fmt.Printf("ch1 2st value: %v, len : %v, cap : %v\n", ch1Received, len(ch1), cap(ch1)) // "", len:0, cap: 2
	// when we receive from closed channel, receiver won't block. we get default value for specific chan type

	//Testing for value sent before closing
	ch1 = make(chan string, 2)
	ch1 <- "Suraj"
	ch1 <- "Hardade"
	close(ch1)
	chRecv1, ok := <-ch1 // Reading second arg from channel indicates value received from channel is sent before
	// channel is closed (true) or after the the channel has closed (false)
	fmt.Printf("ch1 1st value: %v, ok : %v\n", chRecv1, ok) // Suraj, true
	chRecv1, ok = <-ch1
	fmt.Printf("ch1 1st value: %v, ok : %v\n", chRecv1, ok) // Hardade, true
	chRecv1, ok = <-ch1
	fmt.Printf("ch1 1st value: %v, ok : %v\n", chRecv1, ok) //  , false

	// Testing for reading value from empty channel without closing ==> causes deadlock
	ch1 = make(chan string, 2)
	ch1 <- "Suraj"
	fmt.Printf("!st value: %v, len: %v, cap: %v\n", <-ch1, len(ch1), cap(ch1))
	// fmt.Printf("2nd value: %v, len: %v, cap: %v\n", <-ch1, len(ch1), cap(ch1)) // causes deadlock as no go routine send value

	//Incorrect way's of iterating over a channel's values
	// fillCh(5, 5) // for full channel ==> works
	// for i := 0; i < cap(ch); i++ {
	// 	fmt.Println(<-ch) // deadlocak
	// }

	// fillCh(5, 3) // will not works  ==> causes deaklock while reading from channel for 4th value
	// for i := 0; i < cap(ch); i++ {
	// 	fmt.Println(<-ch)
	// }

	// fillCh(5, 5) // full channel ==> will not work as expected as len of channel get updated wvery time when we read data from channel
	// for i := 0; i < len(ch); i++ {
	// 	fmt.Println(<-ch) // 0,1,2  ==> but channel is fulled and having 0,1,2,3,4 data
	// }

	// fillCh(5, 5)
	// l := len(ch) // l will not modify as we reading it before. But what if another go routine added data to channel
	// for i := 0; i < l; i++ {
	// 	fmt.Println(<-ch)
	// }

	// fillCh(5, 5)
	// for v := range ch {
	// 	fmt.Println(v) //deadlock ==> waiting to receive value from channel
	// }

	// iterating over a channel using the range operator
	// ----
	fillCh(5, 3)
	close(ch) // close channel to tell 'range' no more data is expected
	for v := range ch {
		fmt.Println(v)
	}

	fillCh(5, 1)
	close(ch) // close channel to tell 'range' no more data is expected
	for v := range ch {
		fmt.Println(v)
	}
}

var (
	ch chan int
)

func fillCh(c, l int) {
	ch = make(chan int, c)
	for i := 0; i < l; i++ {
		ch <- i
	}
}
