package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ { // loop
		if i%7 == 0 { // condition and modulo
			fmt.Println(i)
		}
	}
}
