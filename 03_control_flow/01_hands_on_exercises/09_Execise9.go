package main

import "fmt"

var s string = "favSpot"

func main() {
	switch s {
	case "a":
		fmt.Println("Hi")
	case "b":
		fmt.Println("hello")
	case "favSpot":
		fmt.Println("favSpot")
	}
}
