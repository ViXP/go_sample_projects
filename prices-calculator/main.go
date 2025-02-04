package main

import (
	"fmt"

	"example.com/prices-calculator/filemanager"
	"example.com/prices-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}
	processChannel := make([]chan bool, len(taxRates))

	for index, rate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", rate))
		//ui := userinput.New()
		job := prices.NewTaxIncludingPriceProcessingJob(fm, rate)
		processChannel[index] = make(chan bool)
		go job.Process(processChannel[index])
	}

	for _, processor := range processChannel {
		<-processor
	}
}
