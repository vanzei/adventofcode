package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Initialize variables
	var lines []string
	var bitLength int

	// Read all lines
	for scanner.Scan() {
		line := scanner.Text() // Read the current line
		lines = append(lines, line)
		bitLength = len(line) // Assuming all lines have the same length
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Initialize counters for each bit position
	zeroCounts := make([]int, bitLength)
	oneCounts := make([]int, bitLength)

	// Count the occurrences of 0s and 1s for each bit position
	for _, line := range lines {
		for i, bit := range line {
			if bit == '0' {
				zeroCounts[i]++
			} else if bit == '1' {
				oneCounts[i]++
			}
		}
	}

	// Initialize a string builder to build the result string
	var result strings.Builder

	// Build the result string based on the counts
	for i := 0; i < bitLength; i++ {
		if oneCounts[i] > zeroCounts[i] {
			result.WriteString("1")
		} else {
			result.WriteString("0")
		}
	}

	// Convert the result string to an integer
	resultingBinary := result.String()
	resultingInt, err := strconv.ParseInt(resultingBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the complement by subtracting from the maximum value based on the bit length
	maxValue := int64((1 << bitLength) - 1) // This gives us the maximum value for the given bit length
	complementInt := maxValue - resultingInt

	// Convert the complement integer back to a binary string
	complementBinary := strconv.FormatInt(complementInt, 2)

	// Pad the binary string with leading zeros if necessary
	complementBinary = fmt.Sprintf("%0*b", bitLength, complementInt)

	// Print the result string and its complement
	fmt.Println("Resulting binary string:", resultingBinary)
	fmt.Println("Complement binary string:", complementBinary)

	// Convert binary strings to decimal and print
	resultingDecimal := resultingInt
	complementDecimal := complementInt

	fmt.Println("Resulting decimal value:", resultingDecimal)
	fmt.Println("Complement decimal value:", complementDecimal)
	fmt.Println("Result:", complementDecimal*resultingDecimal)
}
