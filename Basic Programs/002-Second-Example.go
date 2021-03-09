package main

import "fmt"

func main() {
	n_bytes,error := fmt.Println("Naveen Kumar kammari",12,23)
	fmt.Println("No.of bytes: ",n_bytes)
	fmt.Println("Error code: ",error)
	//Go won't allow unused variables
	// _ will be used for unused variables
	// n,_:= fmt.Println("Naveen",12,true)
	//fmt.Println(n)
	//Here we can't worry about the _ variable
}