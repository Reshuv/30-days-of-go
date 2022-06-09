package main

import "fmt"

func main() {
	var a [5]int // array is collection of similar of single type of elements
	var b [6]int // array type is different because of it's len

	a[3] = 42
	fmt.Println(a)
	fmt.Println(len(a)) // len of an array
	fmt.Println(b)

}
