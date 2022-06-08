// condition you can use if statement or switch statement

package main

import "fmt"

func main() {
	n := "GO"
	switch n {
	case "Hi":
		fmt.Println("learning go")
	case "GO":
		fmt.Println("Easy to learn")
		fallthrough // prints the next statement, if the previous one is correct
	case "no":
		fmt.Println("This is an error")
	default:
		fmt.Println("This is default")
	}
}
