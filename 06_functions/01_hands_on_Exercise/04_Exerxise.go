package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() { //func person having speak method
	fmt.Println("Hi I am", p.first, p.last)
}

func main() {

	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p.speak() //calling method

}
