package main

import "fmt"

type (
	Person struct {
		firstName, lastName string
		Age                 // embedded/Anonymous field
		Address             // embedded/Anonymous field
	}

	Age struct {
		age uint8
	}

	Address struct {
		string // embedded/Anonymous field ==> fileds having only type and empty fieldname
		// string  ==> we can't have more than one anonymous field having same type ==> causes ambugity
		dist string
		pin  string
	}
)

func main() {

	//initialization of nasted structure
	p1 := Person{
		firstName: "Suraj",
		lastName:  "Hardade",
		Age:       Age{10},
		Address:   Address{"Madha", "solapur", "410421"},
	}

	fmt.Printf("%#v\n", p1)
	fmt.Printf("Age : %d, Address: %v, %v, %v", p1.age, p1.string, p1.dist, p1.pin) // field promotion to its uparent struct
	fmt.Printf("Age : %d, Address: %v, %v, %v", p1.Age.age, p1.Address.string, p1.Address.dist, p1.Address.pin)

	p2 := Person{
		firstName: "Krut",
		lastName:  "Hardade",
		Age:       Age{age: 10},
		Address:   Address{string: "Madha", dist: "solapur", pin: "410421"},
	}
	fmt.Printf("%#v\n", p2)
}
