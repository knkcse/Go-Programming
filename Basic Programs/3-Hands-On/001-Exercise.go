package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {

		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
		fmt.Println()
		year := 1996
		for year < 2022 {
			fmt.Println(year)
			year++
		}
		for i:=10;i<=100;i++{
			fmt.Println(i%4)
		}
	}
}
