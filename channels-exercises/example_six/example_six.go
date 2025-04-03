package example_six

import (
	"fmt"
	"log"
	"time"
)

func Run() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "from routine"
	}()

	select {
	case from := <-ch:
		fmt.Println(from)
	case <-time.After(5 * time.Second):
		log.Fatal("timeout")
	}
}
