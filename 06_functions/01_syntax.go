//syntax
// func(r receiver) identifier(parameters) (return(s)) {...}
// when you define a func its called a parameter and when you call a func its called argument

package main

import "fmt"

func main() {

	foo()
	bar("Kitty")

	s1 := woo("Penny")
	fmt.Println(s1)

	x, y := mini("KK", "singh")
	fmt.Println(x)
	fmt.Println(y)

}

//basic function

func foo() {
	fmt.Println("hello foo")
}

//everthing in go pass by value

//takes an arguement

func bar(s string) {
	fmt.Println("hello,", s)
}

//return func

func woo(st string) string {
	return fmt.Sprint("Hello,", st)
}

//return multiple func

func mini(fn string, ln string) (string, bool) {
	a := fmt.Sprint("hello,", fn, ln)
	b := false
	return a, b
}
