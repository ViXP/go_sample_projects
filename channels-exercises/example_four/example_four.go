package example_four

import "fmt"

func Run() {
	readChn := createChan()

	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Printf("Received: %v\n", v)
		}
		fmt.Println("Done.")
	}

	consumer(readChn)
}

func createChan() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 6; i++ {
			ch <- i
		}
	}()
	return ch
}
