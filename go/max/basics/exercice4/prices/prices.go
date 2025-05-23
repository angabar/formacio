package prices

import (
	"fmt"

	"example.com/conversion"
	"example.com/iomanager"
)

// Para excluir un valor de un json tenemos que usar "-"
type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tag_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, error := job.IOManager.ReadLines()

	if error != nil {
		fmt.Println(error)
		return error
	}

	prices, error := conversion.StringsToFloat(lines)

	if error != nil {
		fmt.Println(error)
		return error
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	error := job.LoadData()

	if error != nil {
		fmt.Println(error)
		errorChan <- error
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
