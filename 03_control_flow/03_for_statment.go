package main

import "fmt"

func main() {
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")
}

/*Like a C "for"
for init;condition;post{}

Like a C "while"
for condition{}

Like a C for{;;}
for{ }
*/
