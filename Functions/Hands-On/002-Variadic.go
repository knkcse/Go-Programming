package main

import "fmt"

func main() {
	//sum1 := foo(1, 2, 3, 4, 5, 3, 6, 4, 5)//
	x := []int{1, 2, 43, 3, 56, 3}
	sum1 := foo(x...) //unfrolling
	//x := []int{2, 3, 4, 3, 7}
	sum2 := bar([]int{2, 3, 4, 3, 7})
	fmt.Println(sum1)
	fmt.Println(sum2)
}
func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
