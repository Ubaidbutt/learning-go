package main

import "fmt"

func main() {
	var totalRevenue, totalExpenses, taxRate float64

	fmt.Print("Total revenue in USD: ")
	fmt.Scan(&totalRevenue)

	fmt.Print("Total expenses in USD: ")
	fmt.Scan(&totalExpenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := totalRevenue - totalExpenses
	profit := earningBeforeTax * (1 - taxRate/100)

	ratio := earningBeforeTax / profit

	fmt.Printf("Your earnings before tax are: %.2f\n", earningBeforeTax)
	fmt.Printf("Your profit is: %.2f\n", profit)
	fmt.Printf("Your earning before tax to profit ratio is: %.2f\n", ratio)
}
