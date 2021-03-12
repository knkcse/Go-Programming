package main

import "fmt"

func main() {
	fmt.Println("Here we go with the switch case combination")
	switch "A" {
	case "A":
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("#")
		fallthrough
	case "B":
		fmt.Println("This is B case")
	case "C":
		fmt.Println("Naveen Kumar Kammari")
	}
}
