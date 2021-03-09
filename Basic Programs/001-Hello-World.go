package main

import "fmt"

func main() {
	fmt.Println("This is my first go programming code!.");
	for i := 0; i < 5; i++ {
		fmt.Println("Naveen Loves coding");
	}
	otherFunction()
	fmt.Println("\n This is after function call");
}
func otherFunction() {
	fmt.Println("\n We are in other Function call. We can use semicolons at the end of the statement as well");
}
