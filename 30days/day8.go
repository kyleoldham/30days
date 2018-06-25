package main

import "fmt"

func main() {
	var array_length int
	fmt.Scanf("%d", &array_length)

	// Create array
	var array = make([]int, array_length)

	// Take in array input
	for i := 0; i < array_length; i++ {
		fmt.Scanf("%d", &array[i])
	}

	// Swippity swap
	for i := 1; i <= len(array); i++ {
		fmt.Printf("%d ", array[array_length-i])
	}
}
