package main

import "fmt"

func main() {
	s1 := make([]int, 10, 100)
	pS2 := new([]int)

	fmt.Printf("s1 : %v, Type: %T, len : %v, cap: %v\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("pS2 : %p, Type: %T, len : %v, cap: %v\n", pS2, pS2, len(*pS2), cap(*pS2))
	fmt.Printf("*pS2 : %v, Type: %T, len : %v, cap: %v\n", *pS2, *pS2, len(*pS2), cap(*pS2))

	fmt.Printf("*pS2 value : %p\n", *pS2)
	s1[0] = 10
	*pS2 = []int{10}
	fmt.Printf("s1 : %v, Type: %T, len : %v, cap: %v\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("pS2 : %v, Type: %T, len : %v, cap: %v\n", *pS2, *pS2, len(*pS2), cap(*pS2))

	s1 = append(s1, 110)
	fmt.Printf("s1 : %v, Type: %T, len : %v, cap: %v\n", s1, s1, len(s1), cap(s1))

	var ptest = new(int)
	var x int

	fmt.Printf("ptest: %v, Type: %T, *ptest: %v, *ptest Type:%T\n", ptest, ptest, *ptest, *ptest)
	fmt.Printf("x : %v, Type: %T\n", x, x)
}
