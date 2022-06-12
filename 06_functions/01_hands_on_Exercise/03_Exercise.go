package main

import "fmt"

func main() {

	defer one()
	two()
	three()
	defer four()

}

func one() {
	defer func() {
		fmt.Println("Sub One")
	}()
	fmt.Println("ONE")
}

func two() {
	defer func() {
		fmt.Println("Sub two")
	}()
	fmt.Println("TWO")
}

func three() {
	fmt.Println("THREE")
}

func four() {
	defer func() {
		fmt.Println("Sub four")
	}()
	fmt.Println("FOUR")
}
