package main

import (
	"fmt"

	"example.com/cmdmanager"
	"example.com/filemanager"
	"example.com/prices"
)

1:36
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
