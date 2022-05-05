package main

import "fmt"

//decalred a variable of type string holds a value of string, but we didn't assign value
//i.e will print null value

var y string
var z int

func main() {
	fmt.Println("Printing value of y ", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z) //print the value of int i.e. zero value
	fmt.Printf("%T", z)
}
