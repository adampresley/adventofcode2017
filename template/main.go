package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var err error
	var input []byte
	var reader bytes.Reader

	inputFile := "input.txt"

	if input, err = ioutil.ReadFile(inputFile); err != nil {
		fmt.Printf("Error reading input file: %s\n", err.Error())
		os.Exit(-1)
	}

	reader = bytes.NewReader(input)

	for err == nil {

	}

}
