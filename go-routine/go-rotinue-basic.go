package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-------Go-Routine-------")
	start := time.Now()

	// creating goroutine from named function
	go Producer(1) // go runtime to manages the execution of this func in separate routine/thred
	// Before startig exceution in go routine main thread is terminating hence we will not get any results
	// if we are not waiting in main routine to complete excecution of Procedure routine.
	go Producer(2)
	go Producer(3)

	//creating go routine from ananymous function
	go func() {
		start := time.Now()
		for i := 0; i < 5; i++ {
			fmt.Printf("ananymous function : message # %v\n", i)
		}
		Producer(4)
		fmt.Printf("ananymous function elapsed time : %v\n", time.Now().Sub(start))
	}()
	time.Sleep(1 * time.Microsecond) // waiting for above go-routinr to complete
	fmt.Printf("main thread elapsed time: %v\n", time.Now().Sub(start))
}

func Producer(no int) {
	start := time.Now()
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer #%v- Message : %v\n", no, i)
	}
	elapse := time.Now().Sub(start)
	fmt.Printf("Producer #%v : elapsed time : %v\n", no, elapse)
}
