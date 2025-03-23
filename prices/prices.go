package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/fileManager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData(filePath string) {
	lines, err := fileManager.ReadLines(filePath)

	if err != nil {
		fmt.Println(err)
		return;
	}

	prices, err := conversion.StringsToFloat(*lines)
	if err != nil {
		fmt.Println(err)
		return;
	}

	job.InputPrices = *prices
 }

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData("prices.txt")
	result := make(map[string]string)

	for _, priceVal := range job.InputPrices {
		taxIncludedPrice := priceVal * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", priceVal)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}