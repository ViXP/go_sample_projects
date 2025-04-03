package example_five

import (
	"fmt"
	"time"
)

func Run() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "from two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case one := <-ch1:
			fmt.Println(one)
		case two := <-ch2:
			fmt.Println(two)
		}
	}
}
