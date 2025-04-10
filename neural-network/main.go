package main

import (
	"fmt"

	"example.com/neural-network/neural"
)

func main() {
	neuralNetwork := neural.NewNetwork(1)

	input := float64(10)
	target := float64(input + 2)

	// Training
	train(neuralNetwork, 50, input, target)

	fmt.Printf("Input: %f, Prediction is: %f, Expected is: %f\n", input, neuralNetwork.Forward(input), target)
}

func train(network *neural.Network, iterations int, input float64, target float64) {
	fmt.Println("Training...")
	for i := range iterations {
		network.Train(input, target, 0.001)

		if i%100 == 0 {
			fmt.Printf("Iteration #%d, current prediction is: %f\n", i, network.Forward(input))
		}
	}
}
