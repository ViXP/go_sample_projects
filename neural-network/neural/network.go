package neural

type Network struct {
	hiddenLayer []*Neuron
	outputLayer *Neuron
}

func NewNetwork(hiddenSize int) *Network {
	hiddenLayer := make([]*Neuron, hiddenSize)

	for i := range hiddenSize {
		hiddenLayer[i] = NewNeuron(1)
	}

	outputLayer := NewNeuron(hiddenSize)

	return &Network{hiddenLayer, outputLayer}
}

func (network *Network) Forward(input float64) float64 {
	hiddenOutput := make([]float64, len(network.hiddenLayer))

	for i, n := range network.hiddenLayer {
		hiddenOutput[i] = n.Activate(input)
	}

	return network.outputLayer.Activate(hiddenOutput...)
}

func (network *Network) Train(input float64, target float64, learnRate float64) {
	firstPass := network.Forward(input)
	predictionError := target - firstPass

	for i := range network.hiddenLayer {
		network.hiddenLayer[i].Bias += predictionError * learnRate
		network.hiddenLayer[i].Weights[0] += predictionError * learnRate * input
	}

	// Update output layer weights
	for i := range network.hiddenLayer {
		network.outputLayer.Weights[i] += predictionError * learnRate * network.hiddenLayer[i].Activate(input)
	}

	network.outputLayer.Bias += predictionError * learnRate
}
