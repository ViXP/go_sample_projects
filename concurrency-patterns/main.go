package main

import (
	"runtime"

	"example.com/concurrency-patterns/cancellation"
	"example.com/concurrency-patterns/fan"
	"example.com/concurrency-patterns/pipeline"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pipeline.Run()
	fan.Run()
	cancellation.Run()
}
