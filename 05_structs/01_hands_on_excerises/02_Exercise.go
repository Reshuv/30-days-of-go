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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, v := range v.ice_cream_flavor {
			fmt.Println(i, v)
		}
		fmt.Println("----------------")
	}

}
