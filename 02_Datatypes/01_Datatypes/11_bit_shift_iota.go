package main

import "fmt"
    
const(  
	_ = iota // iota = 0
    //kb:=1024
	kb= 1 << (iota * 10)// iota (2)=10*10
	mb= 1 << (iota * 10)// iota (4)=100*10
	gb= 1 << (iota * 10)// iota (8)= 1000*10
)


func main() {

	fmt.Printf("%d\t\t%b\n",kb,kb)
	fmt.Printf("%d\t\t%b\n",mb,mb)
	fmt.Printf("%d\t%b\n",gb,gb)
}