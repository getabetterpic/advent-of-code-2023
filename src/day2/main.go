package main

import (
	"advent/day1/src/shared"
	"fmt"
	"strings"
)

type Game struct {
	id      int
	samples Sample
}

type Sample map[string]int

func main() {
	lines, err := shared.ReadLinesFromFile("vendor/day2input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	games := map[int]Game{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		var game_id int
		splitLine := strings.Split(line, ":")
		fmt.Sscanf(splitLine[0], "Game %d", &game_id)
		min_possible_cubes := Game{
			id:      game_id,
			samples: map[string]int{},
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

				if cubes > min_possible_cubes.samples[color] {
					min_possible_cubes.samples[color] = cubes
				}
			}
		}
		games[game_id] = min_possible_cubes
	}

	sum := 0
	for game_id, min_possible_cubes := range games {
		power := min_possible_cubes.samples["red"] * min_possible_cubes.samples["green"] * min_possible_cubes.samples["blue"]
		sum += power
		fmt.Println("Game", game_id, "has power of", power)
	}
	fmt.Println("Total power is", sum)
}
