package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", a)
}
