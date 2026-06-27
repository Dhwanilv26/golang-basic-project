package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, 4)
	// donechans is not a buffered channel, it is a slice that can store 4 bool chans
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices/prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate) // yaha par *pricejob nai use kiya, bcz yaha par * isnt used to store or dereference the pointer, we are just initializing the return value of the function
		// now we can use * to dereference the pointer and do tasks accordingly

		go priceJob.Process(doneChans[index]) // go routines dont return any values, ofc need to use channels for the same
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}

}
