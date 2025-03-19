package example_four

import (
	"fmt"
	"sync"
)

func Run() {
	var wg sync.WaitGroup
	var i uint

	incr := func(wg *sync.WaitGroup) {
		wg.Add(1)

		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value of i is %v\n", i)
		}()
		fmt.Println("return from function")
	}

	incr(&wg)
	incr(&wg)
	incr(&wg)
	wg.Wait()
	fmt.Println("done...")
}
