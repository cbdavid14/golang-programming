package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Number of CPUs", runtime.NumCPU())
	fmt.Println("Number of Goroutines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var mu sync.Mutex
	increment := 0
	limit := 100

	wg.Add(limit)
	for i := 0; i < limit; i++ {
		go func() {
			mu.Lock()
			v := increment
			v++
			increment = v
			//fmt.Println(increment)
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Number of goroutines = ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("The  end value of increment is =", increment)
}
