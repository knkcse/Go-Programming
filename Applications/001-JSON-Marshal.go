package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
}

func main() {
	fmt.Println()
	p1 := person{
		"Naveen",
		"Kammari",
		22,
	}
	p2 := person{
		"Anil",
		"Kumar",
		26,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	var people []person             // people is slice of persons
	people = append(people, p1, p2) //appending to slice
	fmt.Println(people)
	//json.Marshal gonna return bytes of slice and error
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println(string(bs))
	}
}
