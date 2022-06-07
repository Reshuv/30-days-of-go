package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) //false because we dont have assign value to x so by default go assign value "0"
	x = true
	fmt.Println(x) //assigned a value so true

	a := 21
	b := 45
	c := 2
	fmt.Println(a == b) //comparision operator
	fmt.Println(c == c)
	fmt.Println(b > c)
	fmt.Println(b != a)

}
