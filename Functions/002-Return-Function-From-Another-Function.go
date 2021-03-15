package main

import "fmt"

func main() {
	func() { //anonymous function
		fmt.Println("It is anonymous function")
	}()
	func(a, b int) { //anonymous function
		fmt.Println("Sum of the two numbers is ", a+b)
	}(5, 12)

	//Assigning a function to a variable
	f := func() {
		fmt.Println("This is function that is assigned to a variable")
	}
	f() //calling function that is assigned to f variable

	//Return a function from another function

	F := returnFunction() //F will store function that is returned from returnFunction()
	fmt.Println(F())

}

//returnfunction is a function that takes 0 parameters
// and returns a function that returns string
func returnFunction() func() string {
	return func() string {
		return "This is string from return function "
	}
}
