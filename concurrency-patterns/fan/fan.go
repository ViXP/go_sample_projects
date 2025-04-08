package fan

import (
	"fmt"
	"sync"
)

func Run() {
	var chs []<-chan uint64

	numbers := []uint64{10, 20, 5}
	enum := generator(numbers)

	// Fan out
	for range numbers {
		chs = append(chs, factorial(enum))
	}

	// Fan in
	for number := range merge(chs) {
		fmt.Printf("Factorial number: %v\n", number)
	}
}

func merge(chs []<-chan uint64) <-chan uint64 {
	var wg sync.WaitGroup

	out := make(chan uint64)

	for _, ch := range chs {
		wg.Add(1)
		go func() {
			for n := range ch {
				out <- n
				wg.Done()
			}
		}()
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

func generator(numbers []uint64) <-chan uint64 {
	ch := make(chan uint64)

	go func() {
		defer close(ch)
		for _, num := range numbers {
			ch <- num
		}
	}()

	return ch
}

func factorial(in <-chan uint64) <-chan uint64 {
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
			out <- result
		}
	}()

	return out
}
