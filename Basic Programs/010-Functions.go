package main

import "fmt"

func main() {
	fmt.Println("Calling a function without parameters")
	noParam()
	fmt.Println("Calling a function with two parameters ")
	twoParam("Naveen", "Kumar")
	fmt.Println("Calling a function with single return type")
	x := sumTwo(13, 27)
	fmt.Println("Returned value from the function is ", x)
	fmt.Println("Calling a function with multiple return types")
	name, age, msg := multipleFunction("Naveen", "Kumar", 21)
	fmt.Println(name,age,msg)
}
func noParam() {
	fmt.Println("This is function with no parameters")
}
func twoParam(a, b string) {
	fmt.Println("Name is ", a, b)
}
func sumTwo(a, b int) int {
	return a + b

}
func multipleFunction(fn, ln string, age int) (string, int, string) {
	name := fmt.Sprint(fn, ln)
	msg := " Hello! Everyone This is your function "
	return name, age, msg
}
