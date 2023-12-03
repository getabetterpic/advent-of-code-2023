package main

import (
	"advent/day1/src/shared"
	"fmt"
	"regexp"
)

func splitLine(line string) []string {
	// Use regular expression to split the line into array elements
	re := regexp.MustCompile(`(\.{1}|[0-9]|[^a-zA-Z0-9\.])`)
	result := re.FindAllString(line, -1)

	return result
}

func main() {
	lines, err := shared.ReadLinesFromFile("vendor/day3input.txt")
	if err != nil {
		panic(err)
	}
	matrix := make([][]string, 0)
	for _, line := range lines {
		split := splitLine(line)
		if len(split) > 0 {
			fmt.Println(split)
			matrix = append(matrix, split)
		}
	}

	sum := 0
	for rowIndex, row := range matrix {
		for i := 0; i < len(row); i++ {
			element := row[i]
			// if element is a number, check all adjacent cells to see if there is a symbol
			// if there is, add the number to the sum
			isADigit := regexp.MustCompile(`[0-9]`)
			isASymbol := regexp.MustCompile(`[^a-zA-Z0-9\.]+`)
			if isADigit.MatchString(element) {
				cellNumber := findWholeNumber(row, i)

				width := len(cellNumber) + 2

				for j := rowIndex - 1; j < rowIndex+2; j++ {
					added := false
					if j < 0 || j >= len(matrix) {
						continue
					}
					for k := i - 1; k < i+width-1; k++ {
						if k < 0 || k >= len(matrix[j]) {
							continue
						}
						if isASymbol.MatchString(matrix[j][k]) {
							if rowIndex == 4 {
								fmt.Println("rowIndex:", rowIndex, "i:", i, "width:", width)
								fmt.Println("Symbol:", matrix[j][k])
								fmt.Println("Number:", cellNumber)
								fmt.Println("Row:", j)
								fmt.Println("Column:", k)
								fmt.Println()
							}
							sum += shared.StringToInt(cellNumber)
							added = true
							break
						}
					}
					if added {
						break
					}
				}

				i += len(cellNumber)
			}
		}
	}

	fmt.Println("Sum:", sum)
}

func findWholeNumber(line []string, startingIndex int) string {
	// Use regular expression to walk forward in the line and find all contiguous digits
	notADigit := regexp.MustCompile(`[^0-9]`)
	result := line[startingIndex]
	for i := startingIndex + 1; i < len(line); i++ {
		if notADigit.MatchString(line[i]) {
			break
		}
		result += line[i]
	}
	return result
}
