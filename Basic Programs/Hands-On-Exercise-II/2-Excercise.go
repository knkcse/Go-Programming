package main

import "fmt"

const (
	a = 12
	b = "Naveen Kumar"
	c = false
)
const (
	x int    = 34
	y string = "Anil Kumar"
	z bool   = true
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
