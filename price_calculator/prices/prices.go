package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64 // input price converted into string, with corresponding output price in float
}

// job *TaxJob mtlb job is a pointer of type Taxjob, and (*job).taxrate -> used to deference the pointer value

// address se hi sab modify hota hai ,to inside operations -> use & and return *, this is the std pattern for constructors in Go

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	// pointer stores address of something, so returning the address of something
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range (*job).InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}
