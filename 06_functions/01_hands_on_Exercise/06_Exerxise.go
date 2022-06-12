package main

import "fmt"

func main() {
	func() {

		i := 0
		for i = 0; i < 20; i++ {
			fmt.Println("This is an Anonymous function", i)
		}
	}()
}
