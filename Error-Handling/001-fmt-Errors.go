package main

import "fmt"

func main() {
	n, err := fmt.Println("Naveen Kumar")
	if err != nil {
		fmt.Println("Error happened", err)
	}
	fmt.Println(n)
}
