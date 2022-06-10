package main

import "fmt"

func main() {
	//when you want to use struct in only one little area
	p1 := struct { //struct
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "bond",
		age:   32,
	}
	fmt.Println(p1)

	fmt.Println(p1.first, p1.last, p1.age)

}
