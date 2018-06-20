package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	// fmt.println("The total meal cost is " + ((meal_cost * float64(tip_percent/100)) + (meal_cost * float64(tax_percent/100))) - rofl
	// need to all be floats, and float types need to match
	tip := meal_cost * float64(tip_percent) / 100
	tax := meal_cost * float64(tax_percent) / 100
	// add in a grouped var
	add := meal_cost + tip + tax
	// need .0 to reduce extra 0's, and use %f for float type
	fmt.Printf("The total meal cost is %.0f dollars.", add)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	meal_cost, err := strconv.ParseFloat(readLine(reader), 64)
	checkError(err)

	tip_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tip_percent := int32(tip_percentTemp)

	tax_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tax_percent := int32(tax_percentTemp)

	solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
