package main

import "fmt"

var x int
var y string
var z bool
func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	x= 42
	y="Naveen Kumar Kammari"
	z=true
	s:=fmt.Sprintf("%d %s %v",x,y,z)
	fmt.Println(s)
}
