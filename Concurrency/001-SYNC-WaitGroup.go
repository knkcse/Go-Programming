package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Print("Concurrency and parallelism")

	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("ARCH ", runtime.GOARCH)
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("GoRoutines ", runtime.NumGoroutine())
	wg.Add(1) //Kind of synchronization up
	go foo()  //create go routine to execute foo() in parallel
	bar()
	fmt.Println("----------------------------")
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("GoRoutines ", runtime.NumGoroutine())
	wg.Wait()

}
func foo() {
	for i := 1; i < 5; i++ {
		fmt.Println("Here is ,", i)
	}
	wg.Done() //Kind of exiting from cs
}
func bar() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("Even ,", i)
	}
}
