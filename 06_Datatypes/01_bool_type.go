package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 21
	b := 45
	c := 21
	fmt.Println(a == b)
	fmt.Println(c == c)
	fmt.Println(b > c)
	fmt.Println(b != a)

}
