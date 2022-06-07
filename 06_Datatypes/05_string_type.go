package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)//converting string to slice of byte(byte is alias of utf-8)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}
	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

}
