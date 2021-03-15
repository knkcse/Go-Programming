package main

import "fmt"

func main() {
	fmt.Println("Variadic parameters function")
	sum(1,2,3,4,5,3)

}
func sum(x ...int ) {// x ...int can take 0 or more parameters and should be at the end of the params list
	fmt.Println(x)
}
