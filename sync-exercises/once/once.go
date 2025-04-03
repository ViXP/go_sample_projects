package once

import (
	"fmt"
	"sync"
)

func Run() {
	var wg sync.WaitGroup
	var once sync.Once

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			once.Do(initialize)
			fmt.Println("from goroutine")
		}()
	}

	wg.Wait()
}

func initialize() {
	fmt.Println("Executed once")
}
