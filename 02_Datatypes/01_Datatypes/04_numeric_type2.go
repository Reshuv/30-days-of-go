package main

import "fmt"

var x int8 = -128 // int8 = -128 to 127 and uint8(unsigned int) = 0 to 255

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
}
