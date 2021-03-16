package main

import "fmt"

type person struct{
	fname string
	lname string
	age int
}
type human interface {
	speak()
}
func (p *person ) speak() {
	fmt.Println("In person's speak function")
	fmt.Println(p.fname,p.lname)
}

func info(h human){
	fmt.Println(h)

}
func main() {
	p:=person{
		"Naveen Kumar",
		"Kammari",
		22,
	}
	fmt.Println(p)
	info(&p)//as (p *person) speak()
	fmt.Println("Here it is")
	p.speak()

}
