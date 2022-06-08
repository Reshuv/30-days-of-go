package main

import "fmt"

func main() {
	x := 34
	if x == 45 {
		fmt.Println("Our value was 45")
	} else if x == 43 {
		fmt.Println("Our value was 42")
	} else {
		fmt.Println("Our value was not 45")
	}
}
