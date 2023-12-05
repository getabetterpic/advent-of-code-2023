package main

import (
	"advent/day1/src/shared"
	"fmt"
	"strings"
)

func main() {
	lines, err := shared.ReadLinesFromFile("vendor/day4input.txt")
	if err != nil {
		panic(err)
	}

	// each line looks like this: Card   1: 82 41 56 54 18 62 29 55 34 20 | 37 14 10 80 58 11 65 96 90  8 59 32 53 21 98 83 17  9 87 25 71 77 70 73 24
	// we want to split it into two parts, the first part is the card number, the second part is the numbers on the card
	// we want to split the second part into two parts, the first part is the list of winning numbers, the second part is the list of numbers we have

	// create an array the size of the number of lines, each starting at 1

	card_counts := make([]int, len(lines)-1)
	for _, line := range lines {
		// split the line into two parts
		// split the second part into two parts
		// convert the two parts into arrays of ints
		// compare the two arrays of ints
		// print the card number if we have a winning card
		if line == "" {
			continue
		}
		card_id, winning_numbers, my_numbers := parseLine(line)
		winning_numbers_found := find_winning_numbers(winning_numbers, my_numbers)
		if winning_numbers_found > 0 {
			fmt.Println("Card", card_id, "has", winning_numbers_found, "winning numbers")
			for i := 1; i <= winning_numbers_found; i++ {
				card_counts[(card_id-1)+i] += card_counts[card_id-1] + 1
			}
		}
	}
	// sum the array
	sum := len(lines) - 1
	for _, card_count := range card_counts {
		sum += card_count
	}
	fmt.Println("Total scratchcards:", sum)
}

func find_winning_numbers(winning_numbers []int, my_numbers []int) int {
	// loop through each number in my_numbers and count how many of those numbers are in winning_numbers
	winning_numbers_found := 0
	for _, my_number := range my_numbers {
		if contains(winning_numbers, my_number) {
			winning_numbers_found++
		}
	}
	return winning_numbers_found
}

func contains(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

func parseLine(line string) (int, []int, []int) {
	var card_id int
	card_info := strings.Split(line, ":")
	fmt.Sscanf(card_info[0], "Card %d", &card_id)
	winning_numbers, my_numbers := parseCard(card_info[1])
	return card_id, winning_numbers, my_numbers
}

func parseCard(card string) ([]int, []int) {
	card_parts := strings.Split(card, "|")
	winning_numbers := parseNumbers(card_parts[0])
	my_numbers := parseNumbers(card_parts[1])
	return winning_numbers, my_numbers
}

func parseNumbers(numbers string) []int {
	number_strings := strings.Split(numbers, " ")
	number_ints := []int{}
	for _, number_string := range number_strings {
		var number int
		fmt.Sscanf(number_string, "%d", &number)
		if number == 0 {
			continue
		}
		number_ints = append(number_ints, number)
	}
	return number_ints
}
