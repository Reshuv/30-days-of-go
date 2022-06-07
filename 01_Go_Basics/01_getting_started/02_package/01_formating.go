package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	y = 822
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\t%T\n", y, y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\t%T\t", y, y, y, y) //string printing
	fmt.Println(s)
	fmt.Printf("%v", y)

}
