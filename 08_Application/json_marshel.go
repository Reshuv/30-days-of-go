package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "R", //fields in lower case is not taken by marshel
		Last:  "Verma",
		Age:   20,
	}

	p2 := person{
		First: "k",
		Last:  "Sharma",
		Age:   25,
	}

	person := []person{p1, p2}
	fmt.Println(person)

	bs, err := json.Marshal(person) //takes interface and return bytes[] & error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
