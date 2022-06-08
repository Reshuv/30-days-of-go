package main

import "fmt"

func main() {
	bd := 2000

	for {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
