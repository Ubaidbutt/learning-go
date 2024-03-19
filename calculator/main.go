package main

import (
	"fmt"

	"ubaidbutt.com/maths/calculator/division"
	"ubaidbutt.com/maths/calculator/sum"
)

func main() {
	const a = 10
	const b = 50

	fmt.Println("The sum of the two values is: ", sum.Add(a, b))
	fmt.Println("The sum of the two values is: ", division.Divide(b, a))
}
