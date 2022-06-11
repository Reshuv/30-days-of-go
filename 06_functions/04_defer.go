package main

import "fmt"

func main() {

	//defer will defer he execuyion of a funtion until whereever its being called comes to an end.
	defer first()
	second()
	defer third()

}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func third() {
	fmt.Println("third")
}
