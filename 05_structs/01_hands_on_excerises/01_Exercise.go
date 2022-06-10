package main

import "fmt"

type person struct {
	first            string
	last             string
	ice_cream_flavor []string
}

func main() {
	p1 := person{
		first:            "Ram",
		last:             "Singh",
		ice_cream_flavor: []string{"stawberry", "chocolate"},
	}

	p2 := person{
		first:            "Radha",
		last:             "Sharma",
		ice_cream_flavor: []string{"vanilla", "cookie & cream", "chocolate"},
	}

	fmt.Println(p1)
	for i, v := range p1.ice_cream_flavor {
		fmt.Println(i, v)
	}
	fmt.Println(p2)
	for i, v := range p2.ice_cream_flavor {
		fmt.Println(i, v)
	}

	fmt.Println(p1.first, p1.last, p1.ice_cream_flavor)
	fmt.Println(p2.first, p2.last, p2.ice_cream_flavor)

}
