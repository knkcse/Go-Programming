package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter integer value here")
	var x int
	fmt.Scanf("%d",&x)
	var name string
	fmt.Println("Enter your name here")
	fmt.Scanf("%s",&name)
	fmt.Println("Enter some floating value here")
	var f float32
	fmt.Scanf("%f",&f)
	fmt.Println("Enter single character here")
	var ch byte
	fmt.Scanf("%c",&ch)
	fmt.Println("Enter some randome stuff")
	var a int
	var b string
	var c float32
	fmt.Scan(&a,&b,&c)

	fmt.Println("Using Scanf function ")
	fmt.Println(x)
	fmt.Println(name)
	fmt.Println(f)
	fmt.Println("\nUsing single Scan function")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	var P,Q,R string
	var A,B int
	fmt.Println(P)
	fmt.Println(Q)
	fmt.Println(R)
	fmt.Println(A)
	fmt.Println(B)




}
