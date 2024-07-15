package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func filterRating(lines []string, bitLength int, criteria func(int, int) byte) string {
	filteredLines := lines

	for i := 0; i < bitLength; i++ {
		if len(filteredLines) == 1 {
			break
		}

		zeroCount, oneCount := 0, 0
		for _, line := range filteredLines {
			if line[i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}

		selectedBit := criteria(zeroCount, oneCount)
		var newFilteredLines []string
		for _, line := range filteredLines {
			if line[i] == selectedBit {
				newFilteredLines = append(newFilteredLines, line)
			}
		}

		filteredLines = newFilteredLines
	}

	return filteredLines[0]
}

func main() {
	// Open the file
	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Initialize variables
	var lines []string
	var bitLength int

	// Read all lines
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		bitLength = len(line) // Assuming all lines have the same length
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Define the criteria functions
	oxygenCriteria := func(zeroCount, oneCount int) byte {
		if oneCount >= zeroCount {
			return '1'
		}
		return '0'
	}

	co2Criteria := func(zeroCount, oneCount int) byte {
		if zeroCount <= oneCount {
			return '0'
		}
		return '1'
	}

	// Get the oxygen generator rating
	oxygenRatingBinary := filterRating(lines, bitLength, oxygenCriteria)
	oxygenRating, err := strconv.ParseInt(oxygenRatingBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Get the CO2 scrubber rating
	co2RatingBinary := filterRating(lines, bitLength, co2Criteria)
	co2Rating, err := strconv.ParseInt(co2RatingBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Print the results
	fmt.Println("Oxygen generator rating (binary):", oxygenRatingBinary)
	fmt.Println("Oxygen generator rating (decimal):", oxygenRating)
	fmt.Println("CO2 scrubber rating (binary):", co2RatingBinary)
	fmt.Println("CO2 scrubber rating (decimal):", co2Rating)
	fmt.Println("Life support rating:", oxygenRating*co2Rating)
}
