package main

import (
	"fmt"
	"sync"
)



func main() {
	fmt.Println("In main routine")
	var wg sync.WaitGroup
	wg.Add(2)

	go foo(&wg)
	go bar(&wg)
	fmt.Println("This is last statement in the main")
	wg.Wait()
}
func foo(wg *sync.WaitGroup) {
	fmt.Println("We are in foo function")
	fmt.Println("--We are done here--")
	wg.Done()
}
func bar(wg * sync.WaitGroup) {
	fmt.Println("We are in bar function")
	fmt.Println("--We are done here--")
	wg.Done()
}
