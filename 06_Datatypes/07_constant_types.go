package main

import "fmt"

const a = 73

const b = 46.8

const c = "hi" // const(a=3 b=7)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
