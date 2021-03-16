package main

import "fmt"

func main() {
	fmt.Print("This is channel basics")
	ch := make(chan int)
	ch <- 45                           // writing
	fmt.Println("Channel one: ", <-ch) //reading
	//channel will go to infinite times

}
