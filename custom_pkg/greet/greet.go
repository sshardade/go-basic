package greet

import "fmt"

func Greet_msg(name string) { // Greet_msg() is made public by making first char to CAPTIAL and can be called in other packages.
	fmt.Printf("Good Morning, %s!!!\n", name)
}
