package main

import "fmt"

func main() {
	fmt.Println(foo(bar))
}
func bar() {
	fmt.Println("This is bar function")
}
func foo(f func()) string {
	fmt.Println("Before call back")
	f()
	fmt.Println("After call back")
	return "Naveen Kumar"
}
