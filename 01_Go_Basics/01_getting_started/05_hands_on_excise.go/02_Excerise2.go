package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x) //compiler values assign to the variable ae called "ZERO VALUES"
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
