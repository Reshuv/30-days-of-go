// we cant assign pointer to an int but we can assign int to a pointer

package main

import "fmt"

func main() {
	a := 7
	fmt.Println(a)
	fmt.Println(&a) // & gives address of that value //* pointer to where an int is stored
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * derefernce the address // gives value stored at that address
	fmt.Println(*&a)

	*b = 43 // value at that address is 43
	fmt.Println(a)
}
