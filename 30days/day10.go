package main

import "fmt"

// This way almost makes it to easy. TODO: redo later maybe?
func main() {
	var n, count int
	fmt.Scanf("%d", &n)
	// Count number of consecutive ones, using bitwise shift
	for n != 0 {
		n = (n & (n << 1))
		count++
	}
	fmt.Println(count)
}
