package main

import "fmt"

var x int

type kartikey int

var y kartikey

func main() {

	x = 42
	y = 100
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	x = int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
