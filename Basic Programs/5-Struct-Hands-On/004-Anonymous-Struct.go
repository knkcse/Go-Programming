package main

import "fmt"

func main() {
	p := struct {
		first string
		last  string
		age   int
	}{
		first: "Naveen Kumar",
		last:  "Kammari",
		age:   23,
	}
	fmt.Println(p)
}
