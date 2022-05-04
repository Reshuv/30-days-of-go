package main

import "fmt"

var y = 89

func main() {
	//short decalaration operator
	x := 46
	fmt.Println(x)
	//declarationm and assigning a value at same time(of acertain type)
	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("Again:", y)
}
