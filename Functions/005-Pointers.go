package main

import "fmt"

type person struct {
	last, first string
	age         int
}

func main() {
	fmt.Println("Person before changeMe call")
	p := person{
		"Kammari",
		"Naveen Kumar",
		22,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println("Person after changeme call")
	fmt.Println(p)
}
func changeMe(p *person) {
	p.last = "Vadla"
	p.first = " NaveenKumarKammari"
	p.age = 23
}
