package neural

import "math/rand"

type Neuron struct {
	Weights []float64
	Bias    float64
}

func NewNeuron(inputsNum int) *Neuron {
	weights := make([]float64, inputsNum)
	bias := rand.NormFloat64()

	for i := range inputsNum {
		weights[i] = rand.NormFloat64()
	}

	return &Neuron{Weights: weights, Bias: bias}
}

func (neuron *Neuron) Activate(inputs ...float64) float64 {
	sum := float64(0)

	for i, value := range inputs {
		// Linear activation
		sum += value * neuron.Weights[i]
	}

	return sum + neuron.Bias
}
