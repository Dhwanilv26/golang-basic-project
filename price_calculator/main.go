package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, 4)
	errorChans := make([]chan error, 4)
	// donechans is not a buffered channel, it is a slice that can store 4 bool chans
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices/prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate) // yaha par *pricejob nai use kiya, bcz yaha par * isnt used to store or dereference the pointer, we are just initializing the return value of the function
		// now we can use * to dereference the pointer and do tasks accordingly

		go priceJob.Process(doneChans[index], errorChans[index]) // go routines dont return any values, ofc need to use channels for the same
	}
	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// } // cant use the errorchan like this because if there is no error in the code, the main go routine will have to wait till eternity , soln to use "select" statement built specifically for channels

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }

	for index := range taxRates {
		// select statement used to define what channel output we need, the channel output received first wins, rest cases dont get executed here
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("done")
		}
	}

}
