package main

import (
	"fmt"
	"sort"
)

type person struct {
	first, last string
	n           []int
	fav         []string
}

func main() {
	fmt.Println()
	p1 := person{
		"Naveen",
		"Kumar Kammari",
		[]int{4, 6, 2, -5, 7, 2, 9, 4, -8},
		[]string{"H", "r", "R", "anil", "upsc", "ban"},
	}
	fmt.Println(p1)
	p2 := person{
		"Anil Kumar",
		"Kammari",
		[]int{7, 6, 5, 4, 3, -3, -9, -8, -12, 23},
		[]string{"Nav", "Amma", "Bav", "Bunny", "Ammu", "Luck","amma"},
	}
	fmt.Println(p2)
	var people = []person{p1, p2}
	fmt.Println(people)
	for i,p:=range people{
		fmt.Println("Person ",i)
		fmt.Println(p.first,p.last)
		sort.Ints(p.n)
		fmt.Println(p.n)
		sort.Strings(p.fav)
		fmt.Println(p.fav)
		fmt.Println("-----------------------------------")
	}

}
