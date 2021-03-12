package main

import "fmt"

func main() {
	x := 1341
	if checkEven(x) == true {
		fmt.Println("Yes! It is even number")
	} else {
		fmt.Println("No, It is not even number")
	}
}
func checkEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}
