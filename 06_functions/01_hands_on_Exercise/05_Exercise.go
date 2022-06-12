package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) cal() float64 {
	return s.side * s.side

}
func (c circle) cal() float64 {
	return math.Pi * c.radius * c.radius

}
func info(sh shape) { //*
	x := sh.cal()
	fmt.Println(x)
}

type shape interface {
	cal() float64 //*
}

func main() {
	s := square{
		side: 4,
	}
	fmt.Println("side=", s)

	c := circle{
		radius: 3,
	}
	fmt.Println("radius:", c)

	info(c)
	info(s)
}
