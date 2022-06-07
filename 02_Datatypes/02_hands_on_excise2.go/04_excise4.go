package main

import "fmt"

var a int = 34

func main(){

	fmt.Printf("%d\t\t%b\t\t%#x\n",a,a,a)
	
	b := a << 1
	fmt.Printf("%d\t\t%b\t\t%#x",b,b,b)

}