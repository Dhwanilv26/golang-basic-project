package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices/prices.txt")
	if err != nil {
		fmt.Println("could not open file:")
		fmt.Println(err)
		return
	}
	defer file.Close() // always close the file

	scanner := bufio.NewScanner(file)
	// getting the file pointer object for the newscanner method by using os.open()
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		// though scanner.scan() returns boolean if error occurs, we can use scanner.error() at the end of the for loop to track and print error
	}
	err = scanner.Err()

	if err != nil {
		fmt.Println("could not read file")
		fmt.Println(err)
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("failed to parse prices in float")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
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
	fmt.Println(result)
}
