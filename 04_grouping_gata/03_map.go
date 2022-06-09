//maps are key-values store
//maps ae unordered list

package main

import "fmt"

func main() {
	m := map[string]int{ // make a map with key and values string is key and int is avlue
		"Jmaes": 32,
		"Joy":   12,
	}
	fmt.Println(m)
	fmt.Println((m["Joy"]))
	fmt.Println(m["Don"]) //enter a key which is not present gives a value 0

	// identifer fo the variable to store this check is OK

	v, ok := m["Don"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Don"]; ok {
		fmt.Println("This is the if Print", v)
	}

	if v, ok := m["Joy"]; ok {
		fmt.Println("This is the if Print", v)
	}

	// add element in map

	m["Todd"] = 45

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 78, 9}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	// delete elemenr in map

}
