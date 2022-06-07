package main

import "fmt"

const(  
	a = iota // iota = 0
	kb= 1 << iota //10
	mb= 1 << iota 
	gb= 1 << iota 
)

func main(){
	fmt.Printf("%d\t\t%b\n",a , a)
	fmt.Printf("%d\t\t%b\n",kb,kb)
	fmt.Printf("%d\t\t%b\n",mb,mb)
	fmt.Printf("%d\t\t%b",gb,gb)
}

