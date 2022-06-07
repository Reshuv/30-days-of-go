package main

import "fmt"

var a = 23
var b = 16
var c = 34

func main ()  {
	fmt.Println(a == b)
	fmt.Println(a >= c)
	fmt.Println(c <= b)
	fmt.Println(c != b)
	fmt.Println(a > c)
	fmt.Println(b < c)
}

