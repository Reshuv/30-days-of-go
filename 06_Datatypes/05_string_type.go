package main

import "fmt"

func main() {
	x := "Hello World"
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	bs := []byte(x)
	fmt.Println(bs)
	fmt.Printf("%T", bs)

}
