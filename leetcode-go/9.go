package main

import "fmt"

func isPalindrome(x int) bool {
	// Can't be palindrome with a negative
	if x < 0 {
		return false
	}

	// Convert to a string
	number_string := strconv.Itoa(x)

	// Check that first and last match, and iterate through to middle
	// Doing two for's on the same line like this is powerful because you can use that i < j for both really easily
	// whereas in two seperate fors, the j wouldn't be declared for the first.
	for i, j := 0, len(number_string)-1; i < j; i, j = i+1, j-1 {
		if number_string[i] != number_string[j] {
			return false
		}
	}
	return true
}
