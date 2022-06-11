package main

import "fmt"

//recurrsion is func() call itself

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)

	n2 := loopfact(5)
	fmt.Println(n2)

}

//recursively

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//loop

func loopfact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
