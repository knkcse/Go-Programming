package main

import "fmt"

type person struct {
	first, last string
	age         int
}
type agentPerson struct {
	person
	ltk bool
}

func (s agentPerson) show() { //attaching show method to agentPerson struct
	fmt.Println(s)
}
func (p person) show() { //attaching show method to person struct
	fmt.Println(p)
}

type human interface {
	show()
}

func bar(h human) {
	fmt.Println("In bar function ", h)
	switch h.(type) {// this checking can help in assertion of type
	case person:
		fmt.Println("The type is person")
		fmt.Println("\tFirst Name:", h.(person).first)//asserting person type in ()
		fmt.Println("\tLast Name:", h.(person).last)
		fmt.Println("\tAge:", h.(person).age)
	case agentPerson:
		fmt.Println("The type is agentPerson")
		fmt.Println("\tFirst Name:", h.(agentPerson).first)
		fmt.Println("\tLast Name:", h.(agentPerson).last)
		fmt.Println("\tAge:", h.(agentPerson).age)
		fmt.Println("\tLike to Kill :", h.(agentPerson).ltk)

	}
}
func main() {
	fmt.Println("This is example for interface and polymorphism")
	p := person{
		first: "Naveen Kumar",
		last:  "Kammari",
		age:   22,
	}
	a := agentPerson{
		person: person{
			first: "Anil Kumar",
			last:  "Kammari",
		},
		ltk: true,
	}
	fmt.Println(p)
	fmt.Println(a)

	//Polymorphism calls
	bar(p)
	bar(a)

}
