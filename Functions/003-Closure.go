package main

import "fmt"

func main() {
	f := clr()
	fmt.Println(f()) //x will be remembered for f() calls
	fmt.Println(f())
	fmt.Println(f())

}
func clr() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
