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
		fmt.Println("line: ", line)
		firstDigit := findFirstDigit(line)
		lastDigit := findLastDigit(line)
		fmt.Println("firstDigit: ", string(firstDigit), " lastDigit: ", string(lastDigit))

		// if firstDigit or lastDigit are empty, skip this line
		if firstDigit == 0 || lastDigit == 0 {
			continue
		}

		// combine the first and last digits into a string and append the results
		combinedDigits := fmt.Sprintf("%s%s", string(firstDigit), string(lastDigit))

		// convert combinedDigits to an integer before appending to results
		combinedInt, err := strconv.Atoi(string(combinedDigits))
		fmt.Println("combinedInt: ", combinedInt)
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
	// first find all word numbers and replace with the digit
	number_words := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var firstIndex = -1
	var digitToReplace string
	for i, _ := range s {
		for word, _ := range number_words {
			if strings.HasPrefix(s[i:], word) {
				digitToReplace = word
				firstIndex = i
				break
			}
		}
		if firstIndex >= 0 {
			break
		}
	}
	s = strings.Replace(s, digitToReplace, number_words[digitToReplace], 1)

	for _, char := range s {
		if unicode.IsDigit(char) {
			return char
		}
	}
	return 0
}

func findLastDigit(s string) rune {
	s = reverse(s)
	reversed_number_words := map[string]string{
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}

	var firstIndex = -1
	var digitToReplace string
	for i, _ := range s {
		if unicode.IsDigit(rune(s[i])) {
			return rune(s[i])
		}
		for word, _ := range reversed_number_words {
			if strings.HasPrefix(s[i:], word) {
				digitToReplace = word
				firstIndex = i
				break
			}
		}
		if firstIndex >= 0 {
			break
		}
	}
	if digitToReplace == "" {
		return 0
	}
	return rune(reversed_number_words[digitToReplace][0])
}
