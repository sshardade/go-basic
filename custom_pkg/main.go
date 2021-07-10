package main

/*
GO111MODULE go env variable is by default enabled from go version 1.13 so to use code using GOPATH variable set  GO111MODULE to off using "go env -w GO111MODULE=off"
go run custom_pkg  ==> 1. can execute cmd from any where. Go finds the package inside $GOROOT/src or $GPATH/src and ran package having main routine.
go install custom_pkg   ==> install binary to $GOPATH/bin if $GOBIN is not set. Else will put binary to $GOBIN dir.
go install custom_pkg/greet   ==> grret package is not having main routine. Hence go will create utility package (greet.a) and placed it in $GOPATH/pkg dir
*/
import (
	"fmt"

	"custom_pkg/greet"
)

func main() {
	name := "Suraj Hardade"
	fmt.Println("Adding pkgs from other resources")
	greet.Greet_msg(name) // Greet_msg() is made public by making first char to CAPTIAL and can be called in other packages.
}
