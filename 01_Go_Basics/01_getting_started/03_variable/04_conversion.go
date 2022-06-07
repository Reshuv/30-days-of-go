package main

import "fmt"

var a int

type hotdog int // created our own type that is  hotdog and underline type is int

var b hotdog

func main() {
	a = 23
	b = 100
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	
	// a=b mot possible because they are of two different type its a static programming lang and dynamically cant be changed
	a = int(b) //converting one type to another, casting is not possible
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
