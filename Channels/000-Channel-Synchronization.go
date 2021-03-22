package main

import (
	"fmt"
	"time"
)

//We can use channels to synchronize execution of the goroutines.
//
func working(ch chan bool) {
	fmt.Println("Naveen Kumar Kammari")
	time.Sleep(time.Second * 2)
	fmt.Println("After sleep, woke up")
	ch <- true
}
func main() {
	ch := make(chan bool, 1)
	go working(ch)
	fmt.Println(<-ch) //Here we will wait untill working function exits
}
