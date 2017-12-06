package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	var err error
	var input []byte
	var currentDigit byte
	var currentDigitInt int

	matchingDigits := make([]int, 0, 5)
	veryFirstDigit := -1
	veryLastDigit := -1
	previousDigit := -1
	currentFirstDigit := -1
	solution := 0

	inputFile := "input.txt"

	if input, err = ioutil.ReadFile(inputFile); err != nil {
		fmt.Printf("Error reading input file: %s\n", err.Error())
		os.Exit(-1)
	}

	reader := bytes.NewReader(input)

	for err == nil {
		currentDigit, err = reader.ReadByte()
		currentDigitInt, _ = strconv.Atoi(string(currentDigit))

		if veryFirstDigit == -1 {
			veryFirstDigit = currentDigitInt
		}

		if err == nil {
			// First digit
			if currentFirstDigit == -1 {
				currentFirstDigit = currentDigitInt
				previousDigit = currentDigitInt
				continue
			}

			// Subsequent digits, with a match to the first
			if currentDigitInt == currentFirstDigit {
				matchingDigits = append(matchingDigits, currentDigitInt)
				previousDigit = currentDigitInt
				continue
			}

			// Subsequent digit, but isn't a match to the first
			if currentDigitInt != currentFirstDigit {
				// If we have a set of matching digits, add to the solution sum
				if len(matchingDigits) > 0 {
					for _, value := range matchingDigits {
						solution += value
					}
				}

				currentFirstDigit = currentDigitInt
				previousDigit = currentDigitInt
				matchingDigits = make([]int, 0, 5)
			}
		}
	}

	if err != io.EOF {
		fmt.Printf("Uh oh. Something happened while reading the file! %s\n", err.Error())
		os.Exit(-1)
	}

	if len(matchingDigits) > 0 {
		for _, value := range matchingDigits {
			solution += value
		}
	}

	// Check the last against the first
	veryLastDigit = previousDigit

	if veryLastDigit == veryFirstDigit {
		solution += veryLastDigit
	}

	fmt.Printf("Solution: %d\n", solution)
}
