package example_one

import "fmt"

func Run() {
	ch := make(chan int)

	go func(a int, b int, ch *chan int) {
		*ch <- a + b
	}(1, 2, &ch)

	c := <-ch

	fmt.Printf("The sum is: %v\n", c)
}
