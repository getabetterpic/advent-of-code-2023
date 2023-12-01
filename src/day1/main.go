package main

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Use os.DirFS to create an fs.FS from the current directory
	fileSystem := os.DirFS(".")

	filePath := "vendor/day1input.txt" // Replace with the path to your file
	lines, err := readLinesFromFile(fileSystem, filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// initialize an array that we can push our results to
	var results []int
	// Print the lines or use them as needed
	for _, line := range lines {
		firstDigit := findFirstDigit(line)
		lastDigit := findFirstDigit(reverse(line))

		// if firstDigit or lastDigit are empty, skip this line
		if firstDigit == 0 || lastDigit == 0 {
			continue
		}

		// combine the first and last digits into a string and append the results
		combinedDigits := fmt.Sprintf("%s%s", string(firstDigit), string(lastDigit))
		if err != nil {
			fmt.Println("Error converting combinedDigits to int:", err)
			return
		}

		// convert combinedDigits to an integer before appending to results
		combinedInt, err := strconv.Atoi(string(combinedDigits))
		if err != nil {
			fmt.Println("Error converting combinedDigits to int:", err)
			return
		}

		results = append(results, combinedInt)
	}

	// now sum the results
	var sum int
	for _, result := range results {
		sum += result
	}
	fmt.Println(sum)
}

func readLinesFromFile(fileSystem fs.FS, filePath string) ([]string, error) {
	// Read the entire file using fs.ReadFile
	fileContent, err := fs.ReadFile(fileSystem, filePath)
	if err != nil {
		return nil, err
	}

	// Convert the byte slice to a string and split into lines
	lines := splitLines(string(fileContent))

	return lines, nil
}

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}

func reverse(s string) string {
	var reversed string
	for _, char := range s {
		reversed = string(char) + reversed
	}
	return reversed
}

func findFirstDigit(s string) rune {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return char
		}
	}
	return 0
}
