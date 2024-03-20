package main

import (
	"fmt"

	"ubaidbutt.com/maths/calculator/division"
	"ubaidbutt.com/maths/calculator/multiplication"
	"ubaidbutt.com/maths/calculator/subtraction"
	"ubaidbutt.com/maths/calculator/sum"
)

func main() {
	var num1, num2 int

	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Println("The sum result produces: ", sum.Add(num1, num2))
	fmt.Println("The division result produces: ", division.Divide(num1, num2))
	fmt.Println("The subtraction result produces: ", subtraction.Subtract(num1, num2))
	fmt.Println("The multiplication result produces : ", multiplication.Multiply(num1, num2))
}
