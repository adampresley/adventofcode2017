package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	var err error
	var input []byte
	var halfDistance int
	var index int
	var currentDigit int
	var halfDistanceDigit int
	var solution int

	inputFile := "input.txt"

	if input, err = ioutil.ReadFile(inputFile); err != nil {
		fmt.Printf("Error reading input file: %s\n", err.Error())
		os.Exit(-1)
	}

	halfDistance = len(input) / 2
	index = 0
	solution = 0

	for err == nil {
		currentDigit, err = getNextDigit(input, index)
		if err == io.EOF {
			continue
		}

		halfDistanceDigit = getHalfDistanceDigit(input, index, halfDistance)

		if currentDigit == halfDistanceDigit {
			solution += currentDigit
		}

		index++
	}

	fmt.Printf("Solution: %d\n", solution)
}

func getNextDigit(input []byte, index int) (int, error) {
	if index > len(input)-1 {
		return 0, io.EOF
	}

	byteValue := input[index]
	result, _ := strconv.Atoi(string(byteValue))

	return result, nil
}

func getHalfDistanceDigit(input []byte, index int, halfDistance int) int {
	location := index + halfDistance

	if location > len(input)-1 {
		location = 0 + (location - len(input))
	}

	byteValue := input[location]
	result, _ := strconv.Atoi(string(byteValue))

	return result
}
