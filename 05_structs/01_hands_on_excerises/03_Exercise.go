package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

func main() {

	type truck struct {
		vehicle
		fourwheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "saffron",
		},
		fourwheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.color)
	fmt.Println(s1.doors)

}
