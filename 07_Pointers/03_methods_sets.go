// method attached to its type is called methid set

package main

import "fmt"

func main() {

	i := info{
		legs: 4,
	}
	zebra1(&i) //values will be only of pointers
}

type info struct {
	legs int
}

type animal interface {
	zebra() int
}

func (i *info) zebra() int { //pointer type of receiver
	return i.legs
}

func zebra1(a animal) {
	fmt.Println(a.zebra())
}
