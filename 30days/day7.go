package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)

	var words []string
	var test_word string

	// Allowing for input of new words, and adding them to the string
	for i := 0; i < num; i++ {
		fmt.Scanf("%s", &test_word)
		words = append(words, test_word)
	}

	// even and odd strings, append the proper characters to the split strings
	for i := range words {
		var even string
		var odd string
		for k := range words[i] {
			if k%2 != 0 {
				odd = odd + string(words[i][k])
			} else {
				even = even + string(words[i][k])
			}
		}
		fmt.Printf("%s %s\n", even, odd)
	}
}
