/*
1. What is channel?..Declaring and Using channels
2. Un-Buffered and Buffered channels
3. Send only and receive only channels
*/

/*
1. What is channel?..Declaring and Using channels
==> Channel provides a mechanism for concurrently running functions to communicate
by sending and receiving values of specific type.
Un-initialized value of channel is nil. */

package main

import "fmt"

func main() {
	//declaring channels
	var ch chan int // ch is of channel int type

	//sending data to channel
	// ch <- 4 // <- send data to channel  ==> error: chan send (nil chan)
	// can't send data to nil channel

	//Receiving from channel
	// <-ch // blocked receiving from nil channel==> error: chan send (nil chan)

	// Un-Buffered channel (creating channel without capacity)
	ch = make(chan int)
	fmt.Printf("ch : %v, ch Type: %T, ch len : %v, ch cap : %v\n", ch, ch, len(ch), cap(ch))

	// sending value on unbuffered channel (without a receiver)
	// ch <- 4 // fatal error: all goroutines are asleep - deadlock! ==> goroutine 1 [chan send]:
	// fails while trying to send.Here channel is valid but no routine is waiting for receive from this channel.
	// when we sending data to unbuffered channels, receiver type must be present on other routine.

	//receiving value from valid unbuffered channel (without a sender)
	// <-ch // fails while trying to receive
	/* fatal error: all goroutines are asleep - deadlock! ===> goroutine 1 [chan receive]:
	un-buffered channels are not usefule when we are having only one go routine.
	*/

	//creating Buffered channels (with capacity)
	ch = make(chan int, 2)
	fmt.Printf("ch : %v, ch Type: %T, ch len : %v, ch cap : %v\n", ch, ch, len(ch), cap(ch))

	//sending data to buffered channel
	ch <- 4
	fmt.Printf("ch : %v, ch Type: %T, ch len : %v, ch cap : %v\n", ch, ch, len(ch), cap(ch))
	/* when we are sending data to buffered channel, we are able to send if there is no receiver is
	present. as we send data to buffered channel length of channel get updated every time we send data.
	*/
	ch <- 8 // also valid as we have channel of capacity 2 // we are able send data to buffered channel till len(ch)==cap(ch)
	fmt.Printf("ch : %v, ch Type: %T, ch len : %v, ch cap : %v\n", ch, ch, len(ch), cap(ch))

	// channel is full for len(ch) == cap(ch)
	//sending data on full buffered channel
	// ch <- 10 // fails  ==> fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan send]:
	// as there is n room on channel to send data on

	// Receiving from buffered channels
	fmt.Printf("First value from ch : %v, len : %v\n", <-ch, len(ch)) // 4, len : 1
	fmt.Printf("First value from ch : %v, len : %v\n", <-ch, len(ch)) // 8, len : 0
	/* Every time if we send data on buffered channel, len of channel gets incremented and when we recieved from it
	len get deremented. we can read from buffered channel till len(ch)  become 0*/

	// Receiving on buffered empty channel
	// fmt.Printf("First value from ch : %v, len : %v\n", <-ch, len(ch))
	/*fatal error: all goroutines are asleep - deadlock! ==>goroutine 1 [chan receive]:
	As no go routine available to send data on buffered channel */

	//send-only and receive only channels
	chs := make(chan string, 10)
	// send only channel
	var out chan<- string // we can make send only channel by adding <- in before type of channel in channel declaration
	out = chs
	out <- "Hello"
	out <- "World"
	fmt.Printf("Send only channel: %T\n", out)

	//Receiving from sedn only channel
	// <-out // invalid operation: <-out (receive from send-only type chan<- string)

	//Receive only channel
	var in <-chan string // receive only channel created by placing '<-' before chan keyword in channel declaration
	in = chs
	fmt.Printf("Receive only channel: type = %T, value: %v\n", in, <-in)
	fmt.Printf("Receive only channel: type = %T, value: %v\n", in, <-in)

	//sending on receive only channel
	//in <- 5 // invalid operation: in <- 5 (send to receive-only type <-chan string)

}
