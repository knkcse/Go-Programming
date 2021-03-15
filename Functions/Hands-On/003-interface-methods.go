package main

import "fmt"

type circle struct {
	r float32
}
type square struct {
	l, w float32
}

func (c circle) area() float32 {
	area := 3.14 * c.r * c.r
	return area
}
func (s square) area() float32 {
	area := s.l * s.w
	return area
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s)
	switch s.(type) {
	case circle:
		fmt.Println("Circle area is ", s.(circle).area())
	case square:
		fmt.Println("Square area is ", s.(square).area())
	}
}
func main() {
	c := circle{
		2.4,
	}
	s := square{
		4,
		5,
	}

	fmt.Println(c)
	fmt.Println(s)
	info(c)
	info(s)

}
