package main

import "fmt"

type person struct {
	first, last string
	favorite    []string
}

func main() {
	mp := map[string]person{}
	p1 := person{
		first:    "Naveen Kumar",
		last:     "Vadla",
		favorite: []string{"Venila", "Chocolate", "Coconut", "Butterscatch", "Amul"},
	}
	p2 := person{
		first:    "Anil Kumar",
		last:     "Kammari",
		favorite: []string{"No", "Mango", "Gova", "Sweet", "Amul", "adsf", "Asdfasdf"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	mp[p1.last] = p1
	mp[p2.last] = p2
	fmt.Println()
	fmt.Println(mp)
	fmt.Println()
	for k,p:=range mp{
		fmt.Println("For key: ",k)
		for i,v:=range p.favorite{
			fmt.Println(i,v)
		}
	}
}
