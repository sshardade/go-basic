// Section 08 - Lecture 02 : Creating Pointers
package main

import "fmt"

func main() {
	// create a pointer using new() on custom types
	// ----
	type ID string
	pId := new(ID)
	fmt.Printf("pId's value: %v, type: %T, *pId:%v\n", pId, pId, *pId)

	type Student struct {
		name string
		age  uint8
		ssn  ID
	}
	pStudent := new(Student)
	fmt.Printf("pStudent's value: %v, address: %p, type: %T, *pStudent: %v\n", pStudent, pStudent, pStudent, *pStudent)
}
