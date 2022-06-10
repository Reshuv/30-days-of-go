package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:       []string{`shaken`, `not stirred`},
		`robbionson_penny`: []string{`writting`, `eating`, `coding`},
		`singh_rohit`:      []string{`travelling`, `swimming`, `treking`},
	}
	fmt.Println(m)

	m["bosco_don"] = []string{`computer`, `data analysis`, `learning`}

	fmt.Println(m)

	delete(m, "singh_rohit")
	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i)
		for j, v2 := range v {
			fmt.Println("\t", j, v2)
		}
	}

}
