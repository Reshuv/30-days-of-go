package main

import "fmt"

func main() {

	/*if !true {
		fmt.Println(false)
	}
	if !false {
		fmt.Println(true)
	}
	if !false {
		fmt.Println("001")
	}
	if !true {
		fmt.Println("002")
	}

	if x := 2; x == 2 {
		fmt.Println(true)
	}

	if x := 2; x != 2 {
		fmt.Println("not equal")
	}

	if !(2 == 2) {
		fmt.Println("2 is not equal")
	}

	if !(2 != 2) {
		fmt.Println("2 is equal")
	}*/

	if x := 42; x == 2 {
		fmt.Println("true")
	}
	fmt.Println("the condition was wrong") // in go we put ";"at the end of every statement,but here we dont do because compiler do it by default
	fmt.Println("here is another statement")
}
