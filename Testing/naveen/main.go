package main

import "fmt"

func main() {
	fmt.Println("5 + 5 +4 ", mySum(5, 5, 4))
	fmt.Println("7 + 5 +4 ", mySum(5, 7, 4))
	fmt.Println("5 + 9 + 6 ", mySum(5, 9, 6))
	fmt.Println("5+7 ", sumTwo(5, 7))

}
func mySum(x ...int) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}
func sumTwo(a, b int) int {
	return a + b
}
