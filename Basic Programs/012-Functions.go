package main

import "fmt"

func main() {
	fmt.Println("Defer")
	defer firstFunction("Naveen", "Kumar") //This will execute after main function exits
	secondFunction()
	//thirdFunction()
}
func firstFunction(a, b string) {
	defer thirdFunction()
	y := fmt.Sprint("Hello! Mr. ", a, b)
	fmt.Println(y)

}
func secondFunction() {
	fmt.Println("This second function but called after defer firstFunction.")
}
func thirdFunction() {
	fmt.Println("This is third function and used with defer call in firstFunction")
}
