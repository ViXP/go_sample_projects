package main

import (
	"example.com/routines-sample/example_one"
	_ "github.com/gorilla/mux"
)

func main() {
	example_one.Run()
	//example_two.Run()
	//example_three.Run()
	//example_four.Run()
	//example_five.Run()
}
