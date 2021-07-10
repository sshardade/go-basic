package main

import "fmt"

type Person struct {
	firstName, lastName string
	age                 uint8
}

func main() {
	p1 := &Person{"p1", "P Family", 25}
	fmt.Printf("p1 value: %v, p1 Type: %T\n", p1, p1)

	fmt.Println("Accessing struct fileds using pointer variable")
	fmt.Printf("%#v, %#v, %d\n", (*p1).firstName, (*p1).lastName, (*p1).age)
	fmt.Printf("%#v, %#v, %d\n", p1.firstName, p1.lastName, p1.age)
}
