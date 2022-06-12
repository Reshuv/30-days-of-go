package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := foo(ii...)
	fmt.Println(s)

	ii2 := []int{12, 45, 12}
	k := bar(ii2)
	fmt.Println(k)
}

//variadic parameter

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

//type int[]

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
