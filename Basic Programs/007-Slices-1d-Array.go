package main

import "fmt"

func main() {
	fmt.Println("Enter your array size")
	var n int
	fmt.Scanf("%d", &n)
	//fmt.Scan(&n)
	arr := []int{}
	var x int
	fmt.Println("Enter your array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	fmt.Println(arr)
	var xy = []int{}
	xy = append(xy, 1, 5, 6, -3, 4)
	fmt.Println()
	fmt.Println(xy)
	xy=append(xy,arr...)// appending two slices
	fmt.Println()
	fmt.Println(xy)
	//s:=[]int{42,43,44,45,46,47,48,49,50}

}
