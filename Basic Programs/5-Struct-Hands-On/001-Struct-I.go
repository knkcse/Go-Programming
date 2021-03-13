package main

import "fmt"

type person struct {
	first, last string
	favorite    []string
}

func main() {
	p1 := person{
		first:    "Naveen Kumar",
		last:     "Kammari",
		favorite: []string{"Venila", "Chocolate", "Coconut", "Butterscatch", "Amul"},
	}
	fmt.Println(p1)
	for i, v := range p1.favorite {
		fmt.Println(i, v)
	}

	p2 := person{
		first:    "Anil Kumar",
		last:     "Kammari",
		favorite: []string{"No", "Mango", "Gova", "Sweet", "Amul","adsf","Asdfasdf"},
	}

	P:=[2]person{}
	P[0]=p1
	P[1]=p2
	fmt.Println()
	for i,x:= range P{
		fmt.Println("record ",i)
		fmt.Println(x.first,x.last)
		for j,v:= range x.favorite{
			fmt.Println("\t",j," ", v)
		}
		fmt.Println()
	}

}
