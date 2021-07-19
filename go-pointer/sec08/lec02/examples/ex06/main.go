// Section 08 - Lecture 02 : Creating Pointers
package main

import "fmt"

func main() {
	// create a pointer using new() on complex types
	// ----
	pArray := new([10]int)
	fmt.Printf("pArray's value: %v, type: %T, *PArray:%v\n", pArray, pArray, *pArray)

	pSlice := new([]float64)
	fmt.Printf("pSlice's value: %v, type: %T, *pSlice:%v\n", pSlice, pSlice, *pSlice)

	pMap := new(map[int64][]complex128)
	fmt.Printf("pMap's value: %v, type: %T, pMap:%v\n", pMap, pMap, *pMap)

	pCh := new(chan *chan string) // => csp:= make(chan *chan string); pCh := &csp
	fmt.Printf("pCh's value: %v, type: %T, *pCh:%v\n", pCh, pCh, *pCh)
}
