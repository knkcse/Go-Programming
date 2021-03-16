package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 34
	}()

	v, ok := <-ch
	fmt.Println(ok, v)
	//If we try to read value from channel again , it will lead to deadlock here
	close(ch) //closing channel
	v, ok = <-ch
	fmt.Println(ok, v)
}
