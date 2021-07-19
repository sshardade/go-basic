//go run .\solution1.go -m 10 -n 4

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	m, n       int
	ch         chan string
	s          = rand.NewSource(time.Now().UnixNano())
	r          = rand.New(s)
	totalCount int
	totalSum   int
)

type prodResult struct {
	count, sum int
}

func init() {
	flag.IntVar(&m, "m", 100, "Max messages per Producer")
	flag.IntVar(&n, "n", 3, "Number of Producers")
	//ch = make(chan string, m*n) always creates capacity based on default values
}

func main() {
	fmt.Println("GoRoutine excercise using sequential method..")

	flag.Parse()
	if n <= 0 || m <= 0 {
		flag.Usage()
	}
	fmt.Printf("Flag values : n = %v, m = %v\n", n, m)

	ch = make(chan string, n*m)
	fmt.Printf("ch : %v, cap : %v\n", ch, cap(ch))

	for id := 1; id <= n; id++ {
		producer(id)
	}
	close(ch)

	consumer()

}

func consumer() {
	prodResultLocal := make(map[string]prodResult)
	for v := range ch {
		//fmt.Printf("Consumer Received : %v\n", v)
		splitRes := strings.Split(v, "- Random Number #")
		//fmt.Println(splitRes)
		prodRandomNum, _ := strconv.ParseInt(splitRes[1], 10, 64)

		pi := prodResultLocal[splitRes[0]]
		pi.count += 1
		pi.sum += int(prodRandomNum)
		prodResultLocal[splitRes[0]] = pi
	}

	for i, v := range prodResultLocal {
		fmt.Println(i)
		fmt.Printf("\tNumber of Elements: %v\n", v.count)
		fmt.Printf("\tSub-total: %v\n", v.sum)
		totalCount += v.count
		totalSum += v.sum
	}

	fmt.Printf("Total count: %v\n\rTotal Sum: %v\n", totalCount, totalSum)
}

func producer(id int) {
	randMsgs := r.Int()%m + 1
	for i := 0; i < randMsgs; i++ {
		s := fmt.Sprintf("Producer #%v - Random Number #%v", id, r.Int()%randMsgs+1)
		ch <- s
	}
}
