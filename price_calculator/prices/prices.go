package prices

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64 // input price converted into string, with corresponding output price in float
}

// there is nothing like p*, always *p
// address se hi sab modify hota hai ,to inside operations -> use & and return *, this is the std pattern for constructors in Go

func newTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
