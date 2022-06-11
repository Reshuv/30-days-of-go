package main

import "fmt"

//callback : callback is passing a func as an argument
//func() can be 1) returned from a func() 2) passed as argument to a func() 3) assigned to a variable

func main() {
	ii := []int{1, 3, 5, 6, 8, 5, 8, 6, 8, 34}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	e := even(sum, ii...)
	fmt.Println("even numbers", e)

	o := odd(sum, ii...)
	fmt.Println("odd numbes", o)
}

func sum(xi ...int) int { //function takes variadic parameter of int and returns an int
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//for even numbers

func even(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}

	}
	return f(y...)
}

//for odd numbers

func odd(f func(xi ...int) int, vi ...int) int {
	var z []int
	for _, v := range vi {
		if v%2 != 0 {
			z = append(z, v)
		}
	}
	return f(z...)
}
