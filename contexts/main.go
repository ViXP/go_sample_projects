package main

import (
	"example.com/contexts/deadline"
	"example.com/contexts/done"
	"example.com/contexts/timeout"
	"example.com/contexts/value"
)

func main() {
	done.Run()
	deadline.Run()
	timeout.Run()
	value.Run()
}
