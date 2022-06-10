//structs is a data structue which allows us to compose together values of different types.
package main

import "fmt"

//composite data structure, aggregate data structure,complex data structue

// we creates a VALUES of certain type that are stored in a VARIABLES
// and those variables have identiferies
// eg:- var a int or type a struct

type person struct { //struct
	first string
	last  string
	age   int
}

func main() {

	p1 := person{ //p1 is composite literal //person is type and then curly braces
		first: "James",
		last:  "bond",
		age:   32,
	}

	p2 := person{ //created a value of type person
		first: "Penny",
		last:  "robbionsion",
		age:   21,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	//embedded struct
	//take one type of struct embedded in another type

	type secretAgent struct {
		person //struct using in anothe struct
		ltk    bool
	}

	sa1 := secretAgent{
		person: person{
			first: "kk",
			last:  "singh",
			age:   32,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "arya",
			last:  "duni",
			age:   12,
		},
		ltk: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(sa2.first, sa2.last, sa2.age, sa2.ltk)

}
