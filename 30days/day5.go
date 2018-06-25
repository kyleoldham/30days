package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)
	for i := 1; i <= 10; i++ {
		sum := input * i
		fmt.Printf("%d x %d = %d\n", input, i, sum)
	}
}
