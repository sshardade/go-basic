// timing out after waiting on a channel

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select statement on Channels: timing out after waiting on a channel..!!")
	ch := producer()
	consume(ch)
}

func consume(ch chan string) {
	for {
		alarm := time.After(1 * time.Millisecond) // modify time duration to see diff results
		select {                                  /* reading from channel using select and range is quite opposite as reading from channel using range
		 will always blocked. It will stop waiting only if we close channel.
		In case of select statement , select won't block/wait on channel. it will select channel if it is ready else execute default case and return .
		*/
		case s, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(s)
		// case s := <-ch:   // without close(ch) this case will work as after some sleep time consumer will return
		// 	fmt.Println(s)
		// case <-time.After(1 * time.Millisecond):
		// 	return
		case <-alarm:
			return
		}
	}
}

func producer() (ch chan string) {
	ch = make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Message #%v from Producer", i)
		}

		time.Sleep(2 * time.Millisecond)

		for i := 11; i < 20; i++ {
			ch <- fmt.Sprintf("Message #%v from Producer", i)
		}

		close(ch) // reading from closed channel is not blocked. If we close channel here then we have to handle it in consume func
	}()
	return
}
