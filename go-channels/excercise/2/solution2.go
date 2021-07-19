package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	m, // max number of messages per producer
	np, // number of producer
	nc int // number of consumer
	ch            chan string
	wgProducer    sync.WaitGroup
	wgConsumer    sync.WaitGroup
	s             = rand.NewSource(time.Now().Unix())
	r             = rand.New(s)
	producerMutex sync.Mutex
	consumerMutex sync.Mutex
	count         int
	totalRecords  map[int]consumerInfo
)

type consumerInfo struct {
	consumerID int
	db         map[string]producerInfo
}

type producerInfo struct {
	message, sum int
}

func init() {
	flag.IntVar(&m, "m", 100, "Maximum number of messages per Producer")
	flag.IntVar(&np, "np", 3, "Number of Producers")
	flag.IntVar(&nc, "nc", 2, "Number of consumers")
	flag.Parse()
}
func main() {
	fmt.Printf("np: %v, nc: %v, m: %v\n", np, nc, m)
	if m <= 0 || nc <= 0 || np <= 0 {
		flag.Usage()
		return
	}

	ch = make(chan string)
	totalRecords = make(map[int]consumerInfo)

	for i := 1; i <= np; i++ {
		producer(i)
	}

	for i := 1; i <= nc; i++ {
		consumer(i)
	}

	wgProducer.Wait()
	close(ch)
	wgConsumer.Wait()

	for _, v := range totalRecords {
		fmt.Printf("Consumer #%v\n", v.consumerID)
		for i, v1 := range v.db {
			fmt.Printf("\t%v\n\t\tNumber of Elements: %v\n\t\tSub-total: %v\n", i, v1.message, v1.sum)
		}
	}

	fmt.Printf("main routine completed successfully : Processed %v messages\n", count)
}

func consumer(id int) {

	wgConsumer.Add(1)
	go func() {
		var record consumerInfo
		db := make(map[string]producerInfo)
		for v := range ch {
			fmt.Printf("Consumer#%v got : %v\n", id, v)
			msgInfo := strings.Split(v, ",")
			randNum, _ := strconv.ParseInt(msgInfo[2], 10, 64)
			tempInfo := db[msgInfo[0]]
			tempInfo.message += 1
			tempInfo.sum += int(randNum)
			db[msgInfo[0]] = tempInfo
		}
		record.consumerID = id
		record.db = db

		fmt.Printf("record: %v\n", record)
		consumerMutex.Lock()
		totalRecords[id-1] = record
		consumerMutex.Unlock()
		wgConsumer.Done()
	}()
}

func producer(id int) {
	wgProducer.Add(1)
	go func() {
		nm := r.Int()%m + 1
		for i := 0; i < nm; i++ {
			ch <- fmt.Sprintf("Producer:%v, Message Random Number,%v", id, r.Int()%m+1)
			producerMutex.Lock()
			count++
			producerMutex.Unlock()
		}
		wgProducer.Done()
	}()
}
