/*
1. Inserting delay
2. using the select keyword
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	//inserting delay the wrong way
	// fmt.Printf("Message 1 at : %v\n", time.Now())
	// sleep(1 * time.Second)
	// fmt.Printf("Message 2 at : %v\n", time.Now())

	//inserting delay with time.sleep()
	// fmt.Printf("Message 1 at : %v\n", time.Now())
	// time.Sleep(1 * time.Second) // pausing the groutine
	// fmt.Printf("Message 2 at : %v\n", time.Now())

	// get notification from channel
	// fmt.Printf("Before : %v\n", time.Now())
	// alarm := notifyAfter(1 * time.Second)
	// <-alarm
	// fmt.Printf("After : %v\n", time.Now())

	//no need to implement to get modified after some sllep. time packges akready provided method for that
	fmt.Printf("Before : %v\n", time.Now())
	<-time.After(1 * time.Second)
	fmt.Printf("After : %v\n", time.Now())

	//select statement : channel selection ==> select uniformally selects channel to send or receive data if multiple channels are ready
	// var ch chan int // nil channel  ==> we can't receive or send or nil channel
	// select {   // will not throw deadlock : if channel is not ready select executes default statement
	// case <-ch:
	// 	fmt.Println("Sending on nil channel ch")
	// default:
	// 	fmt.Println("Can't receive data on nil channel")
	// }

	// select {
	// case <-ch:
	// 	fmt.Println("Sending on nil channel ch")
	// 	// default:
	// 	// 	fmt.Println("Can't receive data on nil channel") // if default commented : fatal error: all goroutines are asleep - deadlock!  ==> chan receive (nil chan)
	// }

	// select {
	// case ch <- 10: // will get deadlock if we were not using select and sending data to nil channel
	// 	fmt.Println("Sending on nil channel ch")
	// default:
	// 	fmt.Println("Can't send data on nil channel") // in case of select statements if all channels are not ready default is executed.
	// }

	// select {
	// case ch <- 10: // will get deadlock if we were not using select and sending/receiving data to nil channel
	// 	fmt.Println("Sending on nil channel ch")
	// case <-ch:
	// 	fmt.Println("Receiving data from nil channel")
	// default:
	// 	fmt.Println("Can't send/receive data on nil channel") // in case of select statements if all channels are not ready default is executed.
	// }

	// sending and receiving from multiple channels
	// var ch1, ch2 chan int
	// select {
	// case ch1 <- 10:
	// 	fmt.Println("Sending on nil channel ch1")
	// case <-ch2:
	// 	fmt.Println("Receiving from nil channel ch2")
	// default:
	// 	fmt.Println("Can't send/receive data on nil channels")
	// }
	const numOfBits = 100

	ch := randomBitsGeneration(numOfBits)
	db := make(map[int8]int)
	fmt.Printf("Random bits: ")
	for v := range ch {
		db[v]++
		fmt.Printf("%v", v)
	}
	fmt.Println()
	fmt.Printf("0 bit: %v\n", db[0])
	fmt.Printf("1 bit: %v\n", db[1])

	for k, v := range db {
		f := (float32(v) / numOfBits) * 100
		fmt.Printf("%v occurred %.2f%% of the time\n", k, f)
	}
	//consume(ch)

}

func consume(ch chan int8) {
	fmt.Printf("Random Bits: ")
	for v := range ch {
		fmt.Printf("%v", v)
	}
	fmt.Println()
}

func randomBitsGeneration(bits int) (ch chan int8) {
	ch = make(chan int8)

	go func() {
		for i := 0; i < bits; i++ {
			select {
			case ch <- 1:
			case ch <- 0:
			}
		}
		close(ch)
	}()
	return
}

func sleep(delay time.Duration) {
	end := time.Now().Add(delay)
	for time.Now().Before(end) {
		//log.Println("Just waiting here....") // still go routine is running by cpu. no other routine get chance to run
	}
}

func notifyAfter(d time.Duration) (alarm chan time.Time) {
	alarm = make(chan time.Time)

	go func() {
		time.Sleep(d)
		alarm <- time.Now()
		close(alarm)
	}()

	return
}
