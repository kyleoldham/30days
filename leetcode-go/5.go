package main

func longestPalindrome(s string) string {
	//EDGECASE:
	//Return string if it's length is 1
	if len(s) < 2 {
		return s
	}

	begin, maxLength := 0, 1

	//Start incrementing through the string
	for i := 0; i < len(s); {
		//EDGECASE:
		// If the remaining characters in the string is less than the largest found palindrome
		// end the checks and return that largest found amount
		if len(s)-i <= maxLength/2 {
			break
		}

		//Two vars both starting at current position i in the string
		//longestPal will increment forwards, oldMatch will check backwards
		oldMatch, longestPal := i, i

		//This increments pal while the count is < stringLen -1 && while s(palVal+1) = s(pal)
		for longestPal < len(s)-1 && s[longestPal+1] == s[longestPal] {
			longestPal++
		}

		//This will make the main for loop way faster by pushing through values that are already shown as matches
		i = longestPal + 1

		//Same intro && oldMatch>0 && s(palVal+1) == s(palVal-1)
		for longestPal < len(s)-1 && oldMatch > 0 && s[longestPal+1] == s[oldMatch-1] {
			longestPal++
			oldMatch--
		}
		//Increment and decrement to check outwards simultaniously for more matches

		newLength := longestPal + 1 - oldMatch

		//Keep track of the largest palindrome value
		if newLength > maxLength {
			begin = oldMatch
			maxLength = newLength
		}
	}

	return s[begin : begin+maxLength]
}

//Extending the palindrome varies in difficulty for odd vs even, need two fors to account for either?
//Take in string input
//check for the largest palendrome
//	loop through string and also loop through the diffent checks for palindrome conditions to be satisfied
// 	Any way to make this better than O(N^2)?
//that will be looping through the string. Maximum length of the string is 1000
//Output the largest palendrome of the string
