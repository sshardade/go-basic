package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	capCh = 10
)

var (
	s          = rand.NewSource(time.Now().UnixNano())
	r          = rand.New(s)
	wgProducer sync.WaitGroup
	wgConsumer sync.WaitGroup
)

func main() {
	fmt.Println("Go-Channels and Go_Routine...!")

	// ch := make(chan string, capCh)
	// producer(1, ch)
	// consumer(ch)

	//concurrency with buffered channel
	// ch := make(chan string, capCh)
	// go producer(1, ch)
	// consumer(ch)

	// *0 Capacity cannel* with concurrent data processor and consumer
	// ch := make(chan string) // unbuffered channel ==> can be only implemented using routine else we get deadlock as channels has no capacity to store msges and hence need both sender and receiver simulteniously
	// go producer(1, ch)
	// consumer(ch)

	//*timed* concurrent data producer ansd consumer
	// ch := make(chan string)
	// go producerTime(1, ch)
	// consumer(ch)

	//multiple concurrent data producer and single consumer
	// ch := make(chan string)
	// producerCon(1, ch)
	// producerCon(2, ch)
	// producerCon(3, ch)
	// go consumer(ch)
	// wgProducer.Wait()
	// close(ch)
	// wgConsumer.Wait()

	//multiple concurrent data producer and two consumer
	ch := make(chan string)
	producerCon(1, ch)
	producerCon(2, ch)
	producerCon(3, ch)
	consumerCon(1, ch)
	consumerCon(2, ch)
	wgProducer.Wait()
	close(ch)
	wgConsumer.Wait()
}

func consumerCon(id int, in chan string) {
	wgConsumer.Add(1)

	go func() {
		db := make(map[string]int)
		for v := range in {
			prod := strings.Split(v, ",")[0]
			db[prod] += 1
		}

		for i, v := range db {
			fmt.Printf("Consumer #%v - Processed %v from #%v\n", id, v, i)
		}
		wgConsumer.Done()
	}()
}

func consumer(in chan string) {
	// consuming all the data sent to the channel
	wgConsumer.Add(1)
	count := 0
	for v := range in {
		count++
		fmt.Printf("Consumer got : %v\n", v) // range is always waiting for receive on open channel
	}

	if count == 0 {
		fmt.Printf("No Data received\n")
		return
	}

	fmt.Printf("Processed %v items\n", count)
	wgConsumer.Done()
}

func producer(id int, out chan string) {
	//create random varible upto capacity of buffered channel(buffered)
	// n := r.Int()%5 /*cap(out)*/ + 1 // cap for unbuffered channel is 0. So we can get panic in the case of unbuffered channels.
	var n int
	if cap(out) == 0 {
		n = r.Int()%10 + 1 // for unbuffered channel
	} else {
		n = r.Int()%cap(out) + 1 // for buffered channel
	}

	fmt.Printf("n: %v\n", n)
	for i := 0; i < n; i++ {
		out <- fmt.Sprintf("Producer #%v, Message #%v", id, i)
	}
	close(out)
}

func producerTime(id int, out chan string) {
	i := 1
	end := time.Now().Add(1000 * time.Millisecond)

	for time.Now().Before(end) {
		out <- fmt.Sprintf("Producer #%v, Message #%v", id, i)
		i++
	}
	close(out)
}

func producerCon(id int, out chan string) {
	wgProducer.Add(1)
	go func() {
		i := 1
		end := time.Now().Add(1000 * time.Millisecond)

		for time.Now().Before(end) {
			out <- fmt.Sprintf("Producer #%v, Message#%v", id, i)
			i++
		}
		wgProducer.Done()
		// close(out) // can't close channel here as other routine also sendingd data to same channel
	}()
}
