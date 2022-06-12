package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	x := 43
	return x
}

// returns int and string

func bar() (int, string) {
	y := 87
	z := "Tea"
	return y, z

}
