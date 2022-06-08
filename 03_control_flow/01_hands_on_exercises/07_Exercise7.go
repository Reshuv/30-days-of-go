package main

import "fmt"

func main() {
	x := 45
	if x == 40 {
		fmt.Println("value is 40")
	} else if x == 44 {
		fmt.Println("value is 44")
	} else {
		fmt.Println("value is 45")
	}
}
