package main

import "fmt"

func main() {
	fmt.Println("hello")
	foo() //its not  anonymous function,it has name

	// anonymous function
	func() { //foo
		fmt.Println("This an Anonymous func")
	}() //()

	func(x int) { //foo
		fmt.Println("This an Anonymous func", x)
	}(42)

	//function expression :- func expression

	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	// another way

	f1 := func(x int) {
		fmt.Println("the year big brother started watching", x)
	}
	f1(1984)

}

func foo() {
	fmt.Println("Go pogramming")
}
