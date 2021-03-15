package main

import "fmt"

func main() {
	a := funcOne()
	name, age := funcTwo()
	fmt.Println("First Function ", a)
	fmt.Println("Second Function: ", name, age)
}
func funcOne() int {
	return 12
}
func funcTwo() (string, int) {
	return "Naveen Kumar", 22
}
