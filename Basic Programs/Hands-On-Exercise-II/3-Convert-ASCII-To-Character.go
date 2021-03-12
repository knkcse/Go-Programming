package main

import "fmt"

func main() {
	fmt.Println("Here we go....")
	for i := 33; i <= 122; i++ {
		//fmt.Printf("ASCII value: %v and String is %s\n", i, string(i))
		fmt.Println()
		fmt.Printf("ASCII value: %v and String is %#U\n", i, i)

	}
}
