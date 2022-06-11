package main

import "fmt"

//return a func() from a func()

func main() {

	//returning string

	s1 := foo()
	fmt.Println(s1)

	//returning func from an anonymous function

	s2 := func() int {
		return 451
	}
	fmt.Printf("%T\n", s2)

	//func from a func()

	x := bar()
	fmt.Println(x())
	//i := x()
	//fmt.Println(i)
	fmt.Printf("%T\n", x)

	//cleaned up

	fmt.Println(bar()())

}

//retruning a string from a func()

func foo() string {
	s := "Returning a string"
	return s
}

//retruning a func() from a func()

func bar() func() int {
	return func() int {
		return 3000
	}
}

// cleaned up

func mini() func() int {
	return func() int {
		return 3000
	}
}
