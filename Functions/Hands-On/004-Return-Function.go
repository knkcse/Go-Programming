package main

import "fmt"

func main() {
	f := foot()
	fmt.Println(f())
}
func foot() func() string {
	return func() string {
		return "Hello, Naveen Kumar"
	}
}
