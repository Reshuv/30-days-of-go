package main

import "fmt"

func main() {
	a := dec()
	b := dec()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func dec() func() int {
	var i int
	return func() int {
		i--
		return i
	}
}
