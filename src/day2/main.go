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
	max_possible_cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// max_actual_cubes := map[string]int{
	// 	"red":   0,
	// 	"green": 0,
	// 	"blue":  0,
	// }

	// create an empty set to push the legit game_ids to
	legit_game_ids := make(map[int]bool)

	for _, line := range lines {
		if line == "" {
			continue
		}
		failed := false
		var game_id int
		splitLine := strings.Split(line, ":")
		fmt.Sscanf(splitLine[0], "Game %d", &game_id)
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
				// check if the number of cubes is greater than the max possible
				if cubes > max_possible_cubes[color] {
					fmt.Println("Game", game_id, "has", cubes, color, "cubes, which is more than the max possible")
					failed = true
					continue
				}
			}
		}
		if failed {
			continue
		}
		legit_game_ids[game_id] = true
	}
	fmt.Println("Legit game ids:", legit_game_ids)
	// sum legit game ids together
	var sum int
	for legit_game_id := range legit_game_ids {
		sum += legit_game_id
	}
	fmt.Println("Sum of legit game ids:", sum)
}
