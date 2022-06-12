package main

import "fmt"

func main() {

	ss := []int{1, 2, 3, 4, 5}
	s := sum(ss...)
	fmt.Println(s)

}

// func as variadic parameter

func sum(f ...int) int {
	sum := 0
	for _, v := range f {
		sum += v
	}
	return sum
}
