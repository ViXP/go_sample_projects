package example_one

import (
	"fmt"
	"time"
)

func Run() {
	fun("direct call")

	go fun("from goroutine")

	go func() { fun("from anonymous") }()

	fv := fun

	go fv("from function value")

	fmt.Println("wait for routines...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done.")
}

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}
