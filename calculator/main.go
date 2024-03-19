package main

import (
	"fmt"

	"ubaidbutt.com/maths/calculator/division"
	"ubaidbutt.com/maths/calculator/multiplication"
	"ubaidbutt.com/maths/calculator/subtraction"
	"ubaidbutt.com/maths/calculator/sum"
)

func main() {
	const a = 10
	const b = 50

	fmt.Println("The sum result produces: ", sum.Add(a, b))
	fmt.Println("The division result produces: ", division.Divide(b, a))
	fmt.Println("The subtraction result produces: ", subtraction.Subtract(b, a))
	fmt.Println("The multiplication result produces : ", multiplication.Multiply(a, b))
}
