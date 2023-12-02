package main

import (
	"advent/day1/src/shared"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Use os.DirFS to create an fs.FS from the current directory
	fileSystem := os.DirFS(".")

	filePath := "vendor/day2input.txt" // Replace with the path to your file
	lines, err := shared.ReadLinesFromFile(fileSystem, filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// max_possible_cubes := map[string]int{
	// 	"red":   12,
	// 	"green": 13,
	// 	"blue":  14,
	// }

	// max_actual_cubes := map[string]int{
	// 	"red":   0,
	// 	"green": 0,
	// 	"blue":  0,
	// }

	games := map[int]map[string]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		var game_id int
		splitLine := strings.Split(line, ":")
		fmt.Sscanf(splitLine[0], "Game %d", &game_id)
		min_possible_cubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		samples := strings.Split(splitLine[1], ";")
		for _, sample := range samples {
			if sample == "" {
				continue
			}
			sample_colors := strings.Split(sample, ",")
			for _, sample_color := range sample_colors {
				if sample_color == "" {
					continue
				}
				var color string
				var cubes int
				fmt.Sscanf(sample_color, "%d %s", &cubes, &color)
				fmt.Println("Game", game_id, "has", cubes, color, "cubes")

				if cubes > min_possible_cubes[color] {
					min_possible_cubes[color] = cubes
				}
			}
		}
		games[game_id] = min_possible_cubes
	}

	sum := 0
	for game_id, min_possible_cubes := range games {
		power := min_possible_cubes["red"] * min_possible_cubes["green"] * min_possible_cubes["blue"]
		sum += power
		fmt.Println("Game", game_id, "has power of", power)
	}
	fmt.Println("Total power is", sum)
}
