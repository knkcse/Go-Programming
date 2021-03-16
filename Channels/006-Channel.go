package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("Value: ", v)
	}

	//using buffer

	ch2 := make(chan int, 100)
	for i := -100; i < 0; i++ {
		ch2 <- i
	}

	close(ch2) // Can we access channel after closing it

	fmt.Println("After closing channel two")
	for v := range ch2 {
		fmt.Println(v)
	}
	//close(ch2)
}
