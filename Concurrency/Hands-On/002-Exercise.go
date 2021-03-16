package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter:=0
	gr:=100

	wg.Add(gr)
	var mu sync.Mutex
	for i:=1;i<=gr;i++{
		go func(){
			mu.Lock()
			value:=counter
			//runtime.Gosched()
			value++
			counter=value
			mu.Unlock()
			wg.Done()
		}()
		//wg.Done()
	}

	wg.Wait()
	fmt.Println("Final counter is ",counter)
	//fmt.Println("After all")
}
