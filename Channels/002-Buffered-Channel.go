package main

import "fmt"

func main() {
	ch := make(chan int, 2) //Channel with buffer 2 (say size)
	ch <- 23
	fmt.Println("Channe one: ", <-ch) //Not going to infinite times here

	ch2 := make(chan int)
	go func() { //This will run in another go routine so channel will be getting value to it
		ch2 <- 34
	}()
	fmt.Println("Channel two: ", <-ch2) // heres waits until something is put in the channel 2
}
