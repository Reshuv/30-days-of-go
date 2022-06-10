package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)

	fmt.Println(a[2:7])
	fmt.Println(a[1:4])
	fmt.Println(a[6:10])

}
