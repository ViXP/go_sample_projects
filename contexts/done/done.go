package done

import (
	"context"
	"fmt"
	"math/rand"
)

func Run() {
	ctx, done := context.WithCancel(context.Background())

	generator := func() <-chan int {
		ch := make(chan int)

		go func() {
			defer close(ch)

			for {
				select {
				case <-ctx.Done():
					fmt.Println("finished")
					return
				case ch <- rand.Int():
				}
			}
		}()
		return ch
	}

	channel := generator()

	i := 1

	for value := range channel {
		if i < 6 {
			fmt.Printf("Generated number is: %d\n", value)
			i++
		} else {
			done()
		}
	}
}
