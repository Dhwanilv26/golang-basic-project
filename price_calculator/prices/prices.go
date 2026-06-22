package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IoManager         filemanager.FileManager // imp: struct ke andar bhi struct define kar sakte
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string // input price converted into string, with corresponding output price in string
}

// job *TaxJob mtlb job is a pointer of type Taxjob, and (*job).taxrate -> used to deference the pointer value

// address se hi sab modify hota hai ,to inside operations -> use & and return *, this is the std pattern for constructors in Go

func NewTaxIncludedPriceJob(fm *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	// pointer stores address of something, so returning the address of something
	return &TaxIncludedPriceJob{
		IoManager:   *fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IoManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("failed to parse prices in float")
		fmt.Println(err)
		return
	}

	job.InputPrices = prices // assigning the prices from the input file to the inputprices field, and ensuring that the job variable is a pointer so that we dont make faaltu copies and save memory
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]string)

	job.LoadData()

	for _, price := range (*job).InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%0.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IoManager.WriteJSON(job)
}
