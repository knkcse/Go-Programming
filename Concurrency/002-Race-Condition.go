package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//fmt.Print()
	count := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			var v int = count
			runtime.Gosched()
			v++
			count = v
			fmt.Println("value v: ", v)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Count is ",count)
}
