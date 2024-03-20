package division

import (
	"fmt"
	"os"
)

func Divide(a, b int) int {
	if b == 0 {
		fmt.Println("Divison by 0 not allowed. Exiting program")
		os.Exit(0)
	}
	return a / b
}
