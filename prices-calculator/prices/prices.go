package prices

import (
	"fmt"
	"strconv"

	"example.com/prices-calculator/conversion"
	"example.com/prices-calculator/iomanager"
)

type TaxIncludingPriceProcessingJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludingPrice map[string]float64  `json:"recalculated_prices"`
}

func NewTaxIncludingPriceProcessingJob(manager iomanager.IOManager, rate float64) *TaxIncludingPriceProcessingJob {
	return &TaxIncludingPriceProcessingJob{
		IOManager: manager,
		TaxRate:   rate,
	}
}

func (job *TaxIncludingPriceProcessingJob) Process(resultChannel chan bool) {
	job.LoadPrices()
	result := make(map[string]float64, len(job.InputPrices))

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = cutToCents(price * (1 + job.TaxRate))
	}

	job.TaxIncludingPrice = result
	job.IOManager.WriteJson(job)
	resultChannel <- true
}

func (job *TaxIncludingPriceProcessingJob) LoadPrices() {
	readPrices, err := job.IOManager.ReadLines()

	if err != nil {
		panic(err)
	}

	prices, err := conversion.StringsToFloats(readPrices)

	if err != nil {
		panic(err)
	}

	job.InputPrices = prices
}

func cutToCents(price float64) (rounded float64) {
	rounded, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", price), 64)
	return
}
