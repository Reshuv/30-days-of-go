package main

import "fmt"

func main() {
	a := [5]int{3, 5, 7, 8, 9} //array
	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i, v)

	}
	fmt.Printf("%T\n", a) //type of array
}
