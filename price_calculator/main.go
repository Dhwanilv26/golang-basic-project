package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {

		priceJob := prices.NewTaxIncludedPriceJob(taxRate) // yaha par *pricejob nai use kiya, bcz yaha par * isnt used to store or dereference the pointer, we are just initializing the return value of the function
		// now we can use * to dereference the pointer and do tasks accordingly
		(*priceJob).Process()
	}

}
