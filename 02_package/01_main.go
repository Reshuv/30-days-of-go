package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello, Playground", 42, true) //_ throws away the return
	fmt.Println(n)
}
