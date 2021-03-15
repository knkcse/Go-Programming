package main

import "fmt"

type person struct {
	first, last string
}

func main() {
	fmt.Println("This is method example")
	p1 := person{
		first: "Naveen Kumar",
		last:  "Kammari",
	}
	fmt.Println(p1)
	p1.show() //Calling show function using p1 person struct
}
func (p person) show() { //func (r receives) name (parameters) (returns){ body}
	//We are attaching show method to person struct
	fmt.Println("First name: ", p.first)
	fmt.Println("Last name: ", p.last)
}
