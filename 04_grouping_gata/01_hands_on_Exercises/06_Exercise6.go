package main

import "fmt"

func main() {

	a := []string{"James", "Bond", "Shaken", "not stirred"}
	fmt.Println(a)

	b := []string{"miss", "Moneypenny", "hloooo", "James"}
	fmt.Println(b)

	c := [][]string{a, b}
	fmt.Println(c)

	for i, d := range c {
		fmt.Println("record:", i)
		for j, v := range d {
			fmt.Printf("\t index position: %v \t the value: \t ", j, v)
		}
	}
}
