package pipeline

import "fmt"

func Run() {
	for number := range square(generator(3, 10, 155)) {
		fmt.Printf("Squared number: %v\n", number)
	}
}

func generator(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num
		}
	}()

	return ch
}

func square(input <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for n := range input {
			ch <- n * n
		}
	}()

	return ch
}
