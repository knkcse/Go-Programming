package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		//fmt.Println("Naveen")
		select {

		case <-q: //if there is anything in channel q return
			fmt.Println("In the case q")
			return //It is printing elements from challen c,
		case v := <-c: //if there is anything in channel c then print
			fmt.Println(v)
		}

	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
