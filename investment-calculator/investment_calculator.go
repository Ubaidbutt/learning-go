package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, expectedReturnOnInvestment float64
	var yearsOfInvestment int
	const inflationRate = 7.5

	fmt.Print("What is the amount you are investing in USD: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("What is the expected return on investment in percentage %: ")
	fmt.Scan(&expectedReturnOnInvestment)

	fmt.Print("How many years are you investing for: ")
	fmt.Scan(&yearsOfInvestment)

	totalAmountAfterInvestment := investmentAmount * math.Pow(1+expectedReturnOnInvestment/100, float64(yearsOfInvestment))
	inflationAdjustedAmount := totalAmountAfterInvestment / math.Pow(1+inflationRate/100, float64(yearsOfInvestment))

	fmt.Println("Your total investment after the total years will be around: ", totalAmountAfterInvestment)
	fmt.Printf("The real amount after adjustment the inflation rate of %v per year will be %v \n", inflationRate, inflationAdjustedAmount)
}
