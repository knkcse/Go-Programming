package main

import "fmt"

func main() {
	fmt.Println()
	x := []string{"Anil", "Naveen", "Kumar", "Kammari", "Laddu", "Kiran", "Varun", "Ammu", "Lucky", "Shreyan", "Bunny"}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i, " ", x[i])
	}
}
