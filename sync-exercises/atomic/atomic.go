package atomic

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func Run() {
	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Your counter is: %v", counter)
}
