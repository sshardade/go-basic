package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName, lastName string
	age                 uint8
}

type Employee struct {
	firstName, lastName string
	Address
}

type Address struct {
	string //city
	pin    string
}

type MyString string

func (s MyString) toUpperCase() string {
	s1 := string(s)
	return strings.ToUpper(s1)
}

/*
func (s string) toUpperCase() string {

}  //invalid receiver string (basic or unnamed type)compiler ==> InvalidRecv
*/

func (a Address) getAddress() string {
	fmt.Printf("..Inside getAddress method by %T receiver type\n", a)
	return a.string + " " + a.pin
}

func (p *Person) modifyFirstName(firstName string) {
	fmt.Println("..Inside modifyFirstName method having pointer type receiver..")
	fmt.Printf("P: %v, &p : %p\n", p, &p)
	p.firstName = firstName // don't have to explicitly dereference p.(handled by go internally)
}

/*
func (p *Person) updateFirstName(firstName string) {
	fmt.Println("..Inside updateFirstName method having value type receiver..")
	fmt.Printf("P: %v, &p : %p\n", p, &p)
	p.firstName = firstName
}
// We can't declare two methods having same name and same type of receiver but having diff value type(value and pointer receiver type)
*/

func (p Person) updateFirstName(firstName string) {
	fmt.Println("..Inside updateFirstName method having value type receiver..")
	fmt.Printf("P: %v, &p : %p\n", p, &p)
	p.firstName = firstName
}

func (p Person) getFullName() string {
	fmt.Println("..Inside getFullName method having value type receiver..")
	fmt.Printf("p: %#v , &p : %p\n", p, &p)
	return p.firstName + " " + p.lastName
}

/*
func getFullName() { // creating func with same name as other but having diff type is not allowed in GO.
	fmt.Println("Another getFullName function with different type\n")
}*/

func getFullName(firstName, lastName string) string {
	fmt.Println("..Inside getFullName function..")
	return firstName + " " + lastName
}

func main() {
	fmt.Println("Methods and Functions in Go...!")
	p := Person{"Suraj", "Hardade", 10}
	fmt.Printf("P = %#v, &p : %p\n", p, &p)

	fmt.Printf("Calling getFullName() function : %v\n", getFullName(p.firstName, p.lastName))
	fmt.Printf("Calling getFullName method : %v\n", p.getFullName()) /* no need to pass arg as method of type recevier
	has access to properties of receiver type*/

	//calling value type reciver to update data
	fmt.Println("---Before calling updateFirstName method --")
	fmt.Printf("P : %v\n", p)
	p.updateFirstName("Krut")
	fmt.Println("---After calling updateFirstName method --")
	fmt.Printf("P : %v\n", p)

	//calling Pointer type reciver to update data
	fmt.Println("---Before calling ModifyFirstName method --")
	fmt.Printf("P : %v\n", p)
	p.modifyFirstName("Krut")
	fmt.Println("---After calling ModifyFirstName method --")
	fmt.Printf("P : %v\n", p)

	//calling above methods (value type and pointer type receiver)
	/* don't affect how we called method of receiver type. Go inetrnally convert it to secific receier type and
	work on it. Hence if we call pointer receiver type method with value of receiver still we will get
	updated results */
	fmt.Println("---Calling updateFistName by providing of address of receiver---")
	(&p).updateFirstName("Shivratna")
	fmt.Printf("p : %v, p : %p\n", p, &p)
	fmt.Println("---Calling ModifyFistName by providing of value of receiver---")
	p.modifyFirstName("Shivratna")
	fmt.Printf("p : %v, p : %p\n", p, &p)

	// Methods for nested structures
	/* we can observe same conecpt of promotion to it parent structure in case of methods also.(like field promotion)*/
	e := Employee{"Suraj", "Hardade", Address{"Pune", "411028"}}
	fmt.Println("--Calling getAddress method using address receiver type--")
	fmt.Printf("Address : %v\n", e.Address.getAddress())
	fmt.Println("--Calling getAddress method using employee receiver type--")
	fmt.Printf("Address : %v\n", e.getAddress()) // promoted method for embedded Address nasted struct field

	/* we can create method for any receiver type pre-defined or usr-defined as long as that type
	defination present inside the same package */
	//E.g : we can create method which is having "string" receiver type as it was defined in another package
	// we can create method for type that we have created(in same package) from "string" type as it having defination in same package
	fmt.Println("--Calling toUpperCase method for MyString type--")
	var s MyString = "Suraj"
	fmt.Printf("Result : %v\n", s.toUpperCase())
}
