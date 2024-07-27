package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// Read all lines
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Get the length of the binary numbers
	bitLength := len(lines[0])

	// Calculate the gamma rate and epsilon rate for part 1
	gammaRate, epsilonRate := calculateGammaEpsilonRates(lines, bitLength)

	// Convert gamma and epsilon rates to decimal
	gammaDecimal, err := strconv.ParseInt(gammaRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilonDecimal, err := strconv.ParseInt(epsilonRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Print the results for part 1
	fmt.Println("Gamma rate:", gammaRate, "Decimal:", gammaDecimal)
	fmt.Println("Epsilon rate:", epsilonRate, "Decimal:", epsilonDecimal)
	fmt.Println("Power consumption:", gammaDecimal*epsilonDecimal)

	// Calculate the oxygen generator rating and CO2 scrubber rating for part 2
	oxygenGeneratorRating := calculateRating(lines, bitLength, true)
	co2ScrubberRating := calculateRating(lines, bitLength, false)

	// Convert oxygen and CO2 ratings to decimal
	oxygenDecimal, err := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2Decimal, err := strconv.ParseInt(co2ScrubberRating, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Print the results for part 2
	fmt.Println("Oxygen generator rating:", oxygenGeneratorRating, "Decimal:", oxygenDecimal)
	fmt.Println("CO2 scrubber rating:", co2ScrubberRating, "Decimal:", co2Decimal)
	fmt.Println("Life support rating:", oxygenDecimal*co2Decimal)
}

// calculateGammaEpsilonRates calculates the gamma and epsilon rates
func calculateGammaEpsilonRates(lines []string, bitLength int) (string, string) {
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

	// Initialize strings for gamma and epsilon rates
	var gammaRate, epsilonRate string

	// Build the gamma and epsilon rates based on the counts
	for i := 0; i < bitLength; i++ {
		if oneCounts[i] > zeroCounts[i] {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	return gammaRate, epsilonRate
}

// calculateRating calculates the rating based on the criteria
func calculateRating(lines []string, bitLength int, mostCommon bool) string {
	candidates := lines
	for i := 0; i < bitLength; i++ {
		if len(candidates) == 1 {
			break
		}
		candidates = filterCandidates(candidates, i, mostCommon)
	}
	return candidates[0]
}

// filterCandidates filters the candidates based on the criteria
func filterCandidates(candidates []string, position int, mostCommon bool) []string {
	var zeroBits, oneBits []string
	for _, candidate := range candidates {
		if candidate[position] == '0' {
			zeroBits = append(zeroBits, candidate)
		} else {
			oneBits = append(oneBits, candidate)
		}
	}
	if mostCommon {
		if len(oneBits) >= len(zeroBits) {
			return oneBits
		}
		return zeroBits
	}
	if len(zeroBits) <= len(oneBits) {
		return zeroBits
	}
	return oneBits
}
