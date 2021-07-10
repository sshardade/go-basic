package main

import "fmt"

func main() {
	fmt.Println("Understanding basics of structure in Go!!!")
	type person struct {
		firstName,
		lastName,
		BY string
		age uint8
	}

	var p2 person
	p2 = person{"P2", "P Family", "NA", 25}
	fmt.Printf("P2 value : %#v, p2 Type: %T\n", p2, p2)

	p3 := person{firstName: "P3", BY: "NA", lastName: "P Family", age: 0}
	//p4 := person{"P4","P Family"} //too few values in struct literal compiler==>InvalidStructLit
	p4 := person{firstName: "p4", lastName: "P Family"} /*valid
	And remaining fileds are initialized to default value */
	/*p5 := person{firstName: "p4", lastName: "P Family", "1996", 25} error:
	mixture of field:value and value initializers*/
	fmt.Printf("p2 : %v, p3 : %v , p4 : %v\n", p2, p3, p4)
}
