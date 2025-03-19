package example_three

import (
	"fmt"
	"sync"
)

func Run() {
	var data int

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		data++
		defer wg.Done()
	}()

	wg.Wait()

	fmt.Printf("the value of data is: %v\n", data)
	fmt.Println("done.")
}
