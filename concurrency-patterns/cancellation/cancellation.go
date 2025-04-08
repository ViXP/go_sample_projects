package cancellation

import (
	"fmt"
	"sync"
)

type doneType chan struct{}

func Run() {
	var chs []<-chan uint64

	done := make(doneType)
	numbers := []uint64{10, 20, 5}
	enum := generator(done, numbers)

	// Fan out
	for range numbers {
		chs = append(chs, factorial(done, enum))
	}

	// Fan in
	for number := range merge(done, chs) {
		fmt.Printf("Factorial number: %v\n", number)

		// Cancellation
		close(done)
	}
}

func merge(done doneType, chs []<-chan uint64) <-chan uint64 {
	var wg sync.WaitGroup

	out := make(chan uint64)

	for _, ch := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range ch {
				select {
				case out <- n:
				case <-done:
					return
				}
			}
		}()
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

func generator(done doneType, numbers []uint64) <-chan uint64 {
	ch := make(chan uint64)

	go func() {
		defer close(ch)
		for _, num := range numbers {
			select {
			case ch <- num:
			case <-done:
				return
			}
		}
	}()

	return ch
}

func factorial(done doneType, in <-chan uint64) <-chan uint64 {
	out := make(chan uint64)

	go func() {
		defer close(out)

		for n := range in {
			result := uint64(1)
			for i := uint64(1); i <= n; i++ {
				result *= i
			}
			if n == 0 {
				result = 1
			}

			select {
			case out <- result:
			case <-done:
				return
			}
		}
	}()

	return out
}
