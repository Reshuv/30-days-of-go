package main

import "fmt"

func main() {

	// use of pointera:-  when you need to change something that is at the certain location
	//1) when you dont want to pass a lot of data
	//2) you want to change data at a location

	// no pointer

	/*x := 0
	//fmt.Println(x)
	foo(x) // the value of this x will be assigned to y
	fmt.Println(x)*/

	//with pointer
	t := 0
	fmt.Println("t before", t)
	fmt.Println("t before", &t)

	bar(&t)
	fmt.Println("t after", t)
	fmt.Println("t after", &t)

}

// no pointer

/*func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}*/

func bar(z *int) {
	fmt.Println("z before", z)
	fmt.Println("z before", *z)

	*z = 43
	fmt.Println("z after", z)
	fmt.Println("z sfter", *z)
}
