package example_two

import "fmt"

func Run() {
	ch := make(chan int, 6)

	go func(cha *chan int) {
		for i := 0; i < 6; i++ {
			fmt.Printf("Sent iteration: %v\n", i)
			*cha <- i
		}
		defer close(*cha)
	}(&ch)

	for iteration := range ch {
		fmt.Printf("Received iteration: %v\n", iteration)
	}
}
