package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

//func (r receiver) identifier(parameter) (return(s)) {code ...}

func (s secretAgent) speak() { // func (r receiver) identifier()
	fmt.Println("HELLO, I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Jammy",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Jimmy",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
