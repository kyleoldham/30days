package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var a uint64
	var b float64
	var c string

	// Read Input
	var wild []string
	for scanner.Scan() {
		wild = append(wild, scanner.Text())
	}

	// Digest Input
	a, _ = strconv.ParseUint(wild[0], 10, 8)
	b, _ = strconv.ParseFloat(wild[1], 1)
	c = wild[2]

	//Outputs
	fmt.Println(i + a)
	fmt.Printf("%.1f\n", d+b)
	fmt.Println(s + c)
}
