package mutex

import (
	"fmt"
	"runtime"
	"sync"
)

func Run() {
	runtime.GOMAXPROCS(4)

	var balance uint
	var wg sync.WaitGroup
	var mx sync.Mutex

	deposit := func(amount uint) {
		defer mx.Unlock()
		mx.Lock()
		balance += amount
	}

	withdraw := func(amount uint) {
		defer mx.Unlock()
		mx.Lock()
		balance -= amount
	}

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}

	wg.Wait()
	fmt.Printf("Your balance: %v", balance)
}
