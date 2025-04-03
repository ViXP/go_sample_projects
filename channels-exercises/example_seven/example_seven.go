package example_seven

import (
	"fmt"
	"time"
)

func Run() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "from routine"
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("did not receive yet")
		}

		fmt.Println("processing...")
		time.Sleep(2 * time.Second)
	}
}
