package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:       []string{`shaken`, `not stirred`},
		`robbionson_penny`: []string{`writting`, `eating`, `coding`},
		`singh_rohit`:      []string{`travelling`, `swimming`, `treking`},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("the record:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
