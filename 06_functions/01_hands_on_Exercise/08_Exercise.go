package main

import "fmt"

func main() {

	f := foo()
	fmt.Println(f())

}

//assigning a function to a variable

func foo() func() int {
	return func() int {
		return 42
	}
}
