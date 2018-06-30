package main

import "fmt"

func main() {
	// First, declare the 6x6 array and fill it with input.
	Array := [6][6]int{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scanf("%d", &Array[i][j])
		}
	}
	// Calulate the hourglass sum for every hourglass in A, and then print the max sum.
	// MaxHGlass, currHGlass -> Compare and swap
	var maxHGlass, currHGlass int
	maxHGlass = -100 // This should be set to a lower amount than the absolute possible min to start?
	// Iterate through first HGlass
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			currHGlass = Array[i][j] + Array[i][j+1] + Array[i][j+2] + Array[i+1][j+1] + Array[i+2][j] + Array[i+2][j+1] + Array[i+2][j+2]
			if currHGlass > maxHGlass {
				maxHGlass = currHGlass
			}
		}
	}
	fmt.Printf("%d", maxHGlass)
}
