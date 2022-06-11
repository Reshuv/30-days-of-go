//intefaces allows us to define behavior and also allows to do polymorphism i.e values can be of one or more type
// a value can be mor then only type
//every type has no methods
package main

import "fmt"

type hotdog int

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() { //interface type bcause of speak method present in interface
	fmt.Println("Hello I am", p.first, p.last, "-Speaks the person")
}

func (s secretAgent) speak() { //interface type because of speak method
	fmt.Println("Hello I am", s.first, s.last, "-Speaks the SecretAgent")
}

type human interface { //interfaces
	speak() //methods ,then its interface
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was pass into bar", h.(person).first) //assertion //assertion is different from conversion
	case secretAgent:
		fmt.Println("I was pass into bar", h.(secretAgent).first) //when you want to go back to concert type of interface
	}
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "jimmy",
			last:  "bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}

	p := person{
		first: "Dr.",
		last:  "Strange",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

	fmt.Println(p)
	p.speak()
	bar(sa1)
	bar(sa2)
	bar(p)

	//conversion //converting a value to a diiferent type
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
