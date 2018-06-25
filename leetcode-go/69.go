package main

import "fmt"

func mySqrt(x int) int {
	result := x

	// ensure non-negative int
	for result*result > x {
		// formula, and it auto truncates, because using ints
		result = (result + x/result) / 2
	}

	return result
}
