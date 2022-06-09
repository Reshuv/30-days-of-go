//grouping data:- taking values of similar type or different type and putting them into one data structure.
//grouping data :- type: 1) slice  2)map

package main

import "fmt"

func main() {

	// a SLICE allows you to group together VALUES of same TYPE

	//COMPOSITE LITERALS = it creates values for structs,arrays and maps and create a new values each time whic are evaluated.
	// x:=type{values} // composite literal

	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0]) //gives index values
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	// SLICE- for range = range clause can manage over the loops in arrays, maps,structs.

	for i, v := range x { // i=index,v=values //range over x is slice
		fmt.Println(i, v) // gives index and values
	}

	// 2nd way

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	//SLICE= slicing a slice

	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[2:4])

	//SLICE = append a slice // adds elements to the end of the slice

	x = append(x, 1, 2, 3)
	fmt.Println(x)

	y := []int{234, 124, 563, 765, 234}
	x = append(x, y...) // ... means can add unlimited values
	fmt.Println(x)

	//SLICE = deleting from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	//SLICE = make // creating a slice a keyword using a key word using make // make([]Type,length,capacity)
	// make is used to create a new array is the value in array exceeds and throws away the old array

	z := make([]int, 10, 12)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z[1] = 14
	z[7] = 45

	z = append(z, 45)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, 90)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	// creates a  new as value in slice exceeds, its doubled the size of array

	z = append(z, 67)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	//SLICE = multi-dimensional-slice

	od := []string{"Go", "Java", "Python", "c++", "c"}
	fmt.Println(od)

	td := []string{"easy", "independent", "opp", "dependent", "beggining"}
	fmt.Println(td)

	md := [][]string{od, td}
	fmt.Println(md)

}
