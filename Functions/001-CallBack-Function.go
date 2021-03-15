package main

import "fmt"

func main() {
	fmt.Println("Here we go")
	xi := []int{1, 3, 5, 2, 6, 8, 3, 4, 6, 10, 23} //slice of int
	s := sum(xi...)
	fmt.Println("Total SUM is ", s)
	fmt.Println("Even numbers sum is ", even(sum, xi...))
}
func sum(xi ...int) int {
	var total int
	for _, v := range xi {
		total += v
	}
	return total
}
func even(f func(xi ...int) int, vi ...int) int {//taking function as argument so callback
	var evenNumbers []int
	for _, v := range vi {
		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		}

	}
	return f(evenNumbers...)
}
