package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	var bitLength int

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		bitLength = len(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		
	}

	zeroCounts := make([]int, bitLength)
	oneCounts := make([]int, bitLength)

	for _, line := range lines {
		for i, bit := range line {
			if bit == '0' {
				zeroCounts[i]++
			} else if bit == '1' {
				oneCounts[i]++
			}
		}
	}

	var result strings.Builder

	for i := 0; i < bitLength; i++ {
		if oneCounts[i] > zeroCounts[i] {
			result.WriteString("1")
		} else {
			result.WriteString("0")
		}
	}
	resultingBinary := result.String()
	resultingDecimal, err := strconv.ParseInt(resultingBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	