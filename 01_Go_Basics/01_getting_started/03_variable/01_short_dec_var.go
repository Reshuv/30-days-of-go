package main

import "fmt"

var y = 89 // used gobally

func main() {
	//short declaration operator
	//declaration and assigning a value at same time(of acertain type)
	x := 46  //used locally
	fmt.Println(x)
	// declaring using var keyword
	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("Again:", y)
}
