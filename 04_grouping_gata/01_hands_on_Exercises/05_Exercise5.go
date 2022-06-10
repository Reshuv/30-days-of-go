package main

import "fmt"

func main() {
	m := make([]string, 10, 10)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	states := []string{`uk`, `us`, `new york`, `la`, `america`, `south`, `east`, `west`, `north`, `un`}
	fmt.Println(states)

	for i, v := range states {
		m[i] = v
	}

	fmt.Println(m)
	for i, v := range m {
		fmt.Println(i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i, m[i])
	}
	fmt.Println(len(m))
	fmt.Println(cap(m))
}
