//syntax
// func(r receiver) identifier(parameters) (return(s)) {...}
// when you define a func its called a parameter and when you call a func its called argument

package main

import "fmt"

func main() {

	foo()
	bar("Kitty")

}

func foo() {
	fmt.Println("hello foo")
}

func bar(s string) {
	fmt.Println("hello,", s)
}
