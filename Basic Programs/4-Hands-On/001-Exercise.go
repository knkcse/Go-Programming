package main

import (
	"fmt"
)

func main() {
	var x = [5]int{}
	x[0] = 12
	x[1] = 34
	x[2] = 40
	x[3] = -10
	x[4] = 1
	fmt.Println(x)
	for i,v:= range x{
		fmt.Println("index: ",i," value: ",v)
	}
	fmt.Printf("%T\n",x)
	xy:=[]int{42,43,44,45,46,47,48,49,50,51}
	for i,v:=range xy{
		fmt.Println(i,v)
	}
	fmt.Println(xy[:5])
	fmt.Println(xy[5:])
	fmt.Println(xy[2:7])
	fmt.Println(xy[1:6])

}
