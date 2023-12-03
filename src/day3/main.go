package main

import (
	"advent/day1/src/shared"
	"fmt"
	"regexp"
	"strings"
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
			matrix = append(matrix, split)
		}
	}

	sum := 0
	for rowIndex, row := range matrix {
		for i := 0; i < len(row); i++ {
			element := row[i]
			isADigit := regexp.MustCompile(`[0-9]`)
			isAGear := regexp.MustCompile(`\*`)

			if isAGear.MatchString(element) {
				part_numbers := make([]string, 0)
				for j := rowIndex - 1; j < rowIndex+2; j++ {
					if j < 0 || j >= len(matrix) {
						continue
					}

					for cursor := i - 1; cursor < i+2; cursor++ {
						if cursor < 0 || cursor >= len(matrix[j]) {
							continue
						}
						if isADigit.MatchString(matrix[j][cursor]) {
							numandoffset := findWholeNumber(matrix[j], cursor)
							part_number := numandoffset[:strings.Index(numandoffset, ":")]
							offset := shared.StringToInt(numandoffset[strings.Index(numandoffset, ":")+1:])
							part_numbers = append(part_numbers, part_number)
							if rowIndex == 4 {
								fmt.Println("Symbol:", matrix[j][cursor])
								fmt.Println("Row:", j)
								fmt.Println("Column:", cursor)
								fmt.Println("part_number:", part_number)
								fmt.Println()
							}
							cursor = cursor + offset + len(part_number)
						}
					}
				}
				if len(part_numbers) == 2 {
					fmt.Println(part_numbers[0], "*", part_numbers[1], "=", shared.StringToInt(part_numbers[0])*shared.StringToInt(part_numbers[1]))
					sum += shared.StringToInt(part_numbers[0]) * shared.StringToInt(part_numbers[1])
				}
			}
		}
	}

	fmt.Println("Sum:", sum)
}

func findWholeNumber(line []string, startingIndex int) string {
	// given a line and a starting index, find the rest of the whole number on either side of the starting index
	// return the whole number as a string
	isADigit := regexp.MustCompile(`[0-9]`)
	combined := ""
	offset := 0
	whileLoop := true
	i := startingIndex
	// look left
	for whileLoop {
		if i < 0 {
			whileLoop = false
			continue
		}
		if isADigit.MatchString(line[i]) {
			combined = line[i] + combined
			offset--
			i--
		} else {
			whileLoop = false
		}
	}
	// look right
	whileLoop = true
	i = startingIndex + 1
	for whileLoop {
		if i >= len(line) {
			whileLoop = false
			continue
		}
		if isADigit.MatchString(line[i]) {
			combined = combined + line[i]
			offset++
			i++
		} else {
			whileLoop = false
		}
	}
	offset = offset + 1
	fmt.Println("offset:", offset)
	return combined + ":" + fmt.Sprint(offset)
}
