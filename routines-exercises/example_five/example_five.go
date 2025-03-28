package example_five

import (
	"fmt"
	"sync"
)

func Run() {
	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			fmt.Printf("Current value is %v\n", index)
		}(i)
	}
	wg.Wait()
}
