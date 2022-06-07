package main

import "fmt"

func main() {
	fmt.Println("I am leaning go and this is control flow with go")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	foo()

	fmt.Println("Hi I am Learning go")
}

func foo() {
	fmt.Println("Control flow with go")
}
