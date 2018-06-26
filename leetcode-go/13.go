package main

import "fmt"

func romanToInt(s string) int {

	// Needs to be byte for comparability
	roman_map := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	end := 0
	total := 0

	// Needs to go backwards b/c roman numerals
	for i := len(s) - 1; i >= 0; i-- {
		temp := roman_map[s[i]]

		sign := 1

		if temp < end {
			sign = -1
		}

		total += sign * temp
		end = temp
	}
	return total
}
