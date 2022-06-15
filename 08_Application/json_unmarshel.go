package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"R","Last":"Verma","Age":20},{"First":"k","Last":"Sharma","Age":25}]` // slice because its n array in json
	bs := []byte(s)
	fmt.Println(s)
	//fmt.Println(bs)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{}

	var people = []person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
