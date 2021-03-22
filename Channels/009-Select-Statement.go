package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		fmt.Println("In first routine...")
		time.Sleep(time.Second)
		ch1 <- "Naveen"
	}()
	go func() {
		fmt.Println("In second goroutine...")
		time.Sleep(time.Second * 2)
		ch2 <- "Anil"
	}()
	//fmt.Println(<-ch1)
	//fmt.Println(<-ch2)
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
