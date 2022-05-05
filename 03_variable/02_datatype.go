package main

import "fmt"

//Declared variable with identifier z is of type string
var y = 42
var z = "strange"

//back ticks

var a string = `James said, "Shaken, not stirred`

func main() {

	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println(z)
	fmt.Printf("%T", z)
	fmt.Println(a)
	fmt.Printf("%T", a)

}
