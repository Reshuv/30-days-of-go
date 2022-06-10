package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{

		first: "Ram",
		friends: map[string]int{"KK": 877, "Rv": 98,
			"TS": 899},
		favDrinks: []string{"Coke", "Mocktail", "Sost drinks"},
	}
	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.friends)
	fmt.Println(p1.favDrinks)

	for k, v := range p1.friends {
		fmt.Println(k, v)
	}
	for i, v := range p1.favDrinks {
		fmt.Println(i, v)
	}
}
