package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Assigning function to variable")
	}
	f()
	fmt.Printf("%T\n", f)
}
