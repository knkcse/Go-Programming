package main

import "fmt"

var y int

func main() {
	type naveen int
	var x naveen
	fmt.Println("Before Assigning value: ", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println("After assigning value: ", x)
	fmt.Printf("%T\n", x)
	y = (int)(x)
	fmt.Println(y)
}
