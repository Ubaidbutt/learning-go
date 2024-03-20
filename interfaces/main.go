package main

import "fmt"

func printMe(a interface{}) {
	fmt.Println("Value: ", a)
}

func main() {
	const b = 10
	const a = "Ubaid"
	const c = false
	const d = 10.45
	list := []string{"ubaidrbutt", "ubaidrbutt-cool"}

	printMe(a)
	printMe(b)
	printMe(c)
	printMe(d)
	printMe(list)
}
