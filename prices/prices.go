package prices

import (
	"bufio"
	"fmt"
	"os"
	"example.com/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData(filePath string) {
	file, error := os.Open(filePath)

	if error != nil {
		fmt.Println("An error occured opening the file!")
		fmt.Println(error)
		return
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	error = scanner.Err()

	if error != nil {
		fmt.Println("An error occured while reading the file")
		fmt.Println(error)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return;
	}

	job.InputPrices = *prices
	file.Close()
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