package main

import (
	"fmt"
	"time"
)

func naveen(ch chan int) {
	fmt.Println("In the naveen function ")
	time.Sleep(time.Second)
	fmt.Println("Sleeping here.. ")
	time.Sleep(time.Second)
	ch <- 1
}
func anil(ch chan int) {
	fmt.Println("\nIn ANIL function")
	time.Sleep(time.Second)
	fmt.Println("Doing some busy work")
	time.Sleep(time.Second)
	ch <- 2
}
func main() {
	ch := make(chan int, 2)
	go naveen(ch)
	go anil(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
