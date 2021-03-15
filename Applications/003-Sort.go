package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is about sort")
	x := []int{4, 6, 3, -5, 7, -100, 23, 78, -4, -675}
	fmt.Println("Before sort", x)
	sort.Ints(x)
	fmt.Println("After sort", x)
	fmt.Println("\nStrings ")
	str:=[]string{"Naveen ","Kumar","Anil","Kammari"}
	fmt.Println("Strings before sort: ",str)
	sort.Strings(str)
	fmt.Println("Strings after sort: ",str)

}
