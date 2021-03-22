package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}
	fmt.Println("Age")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("Age error: ", err)
	}
	fmt.Println("About to exit")

}
