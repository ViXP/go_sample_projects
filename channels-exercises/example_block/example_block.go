package example_block

import (
	"fmt"
	"sync"
)

func DontRun() {
	var wg sync.WaitGroup

	wg.Add(1)

	go blockingRoutine(&wg)

	wg.Wait()
}

func blockingRoutine(wg *sync.WaitGroup) {
	ch := make(chan int, 4)

	defer close(ch)

	for i := 0; i < 5; i++ {
		ch <- i
	}

	fmt.Println("this will never be shown")
	wg.Done()
}
