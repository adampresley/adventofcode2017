package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func main() {
	var err error
	var input []byte
	var reader *csv.Reader
	var records [][]string

	var checksum int

	inputFile := "input.txt"

	if input, err = ioutil.ReadFile(inputFile); err != nil {
		fmt.Printf("Error reading input file: %s\n", err.Error())
		os.Exit(-1)
	}

	reader = csv.NewReader(bytes.NewReader(input))
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	if records, err = reader.ReadAll(); err != nil {
		fmt.Printf("Error reading input data: %s\n", err.Error())
		os.Exit(-1)
	}

	for lineNumber, record := range records {
		fmt.Printf("Line %d\n----------------------------------------------------\n", lineNumber)
		checksum += getQuotient(record)
	}

	fmt.Printf("Checksum: %d\n", checksum)
}

func getQuotient(record []string) int {
	values := convertToInts(record)

	var left int
	var right int
	var quotient float64
	var winner bool

	winner = false

	for outer, value1 := range values {
		for inner, value2 := range values {
			if inner != outer {
				left = value1
				right = value2

				fmt.Printf("value1 == %d, value2 == %d", value1, value2)

				if right < left {
					left, right = right, left
				}

				quotient = float64(right) / float64(left)
				fmt.Printf(", quotient == %f\n", quotient)

				if quotient == math.Trunc(quotient) {
					fmt.Printf("We have a winner!\n\n")
					winner = true
					break
				}
			}
		}

		if winner {
			break
		}
	}

	return int(quotient)
}

func convertToInts(record []string) []int {
	var value int
	values := make([]int, len(record))

	for index, stringValue := range record {
		value, _ = strconv.Atoi(stringValue)
		values[index] = value
	}

	return values
}
