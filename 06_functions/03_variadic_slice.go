package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("this total", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) //slice of x

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for the item in index posiyion", i, "we are adding value", v, "and the sum is", sum)

	}

	fmt.Println("the total sum is", sum)
	return sum

}
