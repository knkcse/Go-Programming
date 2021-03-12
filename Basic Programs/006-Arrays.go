package main

import "fmt"

func main() {
	fmt.Println("Enter your aray size")
	var n int
	fmt.Scanf("%d", &n)
	var array [10]int
	fmt.Println("Enter your array elements")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Println(array)
}
