package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}
func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
	//close(ch) //We can't close a receiver channel
}
func gen() <-chan int {
	c := make(chan int) //infinitely opening so if we don't make it close then it will lead to deadlock

	go func() { // will take other go routine

		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) //Why here close??
	}()

	return c
}
