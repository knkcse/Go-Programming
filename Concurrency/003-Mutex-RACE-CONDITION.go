package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//fmt.Print()
	fmt.Println("CPUs",runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	count := 0
	var wg sync.WaitGroup
	wg.Add(100)
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			var v int = count

			//runtime.Gosched()
			v++
			count = v
			fmt.Println("value v: ", v)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Count is ", count)
}
