//clousure is used to limit the scope of the vaiable
package main

import "fmt"

var x int //scope of x is everywhere

func main() {

	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)
	{
		y := 71
		fmt.Println(y) //scope of y is limited to this code block
	}

}

func foo() {
	fmt.Println("hi")
	x++
}
