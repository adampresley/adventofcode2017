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

	var biggest int
	var smallest int
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

	for _, record := range records {
		smallest, biggest = getBiggestAndSmallest(record)
		checksum += (biggest - smallest)
	}

	fmt.Printf("Checksum: %d\n", checksum)
}

func getBiggestAndSmallest(record []string) (int, int) {
	var biggest int
	var smallest int
	var value int

	biggest = 0
	smallest = math.MaxInt32

	for _, stringValue := range record {
		value, _ = strconv.Atoi(stringValue)

		if value > biggest {
			biggest = value
		}

		if value < smallest {
			smallest = value
		}
	}

	fmt.Printf("Smallest: %d, Biggest: %d\n", smallest, biggest)
	return smallest, biggest
}
