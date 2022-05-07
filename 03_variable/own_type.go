package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 24
	b = 67
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T", b)

}
