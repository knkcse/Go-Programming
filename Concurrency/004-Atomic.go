package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	//fmt.Print()
	fmt.Println("CPUs",runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	var count int64
	var wg sync.WaitGroup
	wg.Add(100)
	//var mu sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			//mu.Lock()
			atomic.AddInt64(&count,1)
			runtime.Gosched()
			fmt.Println("Counter-: ",atomic.LoadInt64(&count))

			//fmt.Println("value v: ", v)
			//mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GO ROUTINES ",runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Final Count is ", count)
}
