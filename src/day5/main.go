package main

import (
	"advent/day1/src/shared"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := shared.ReadLinesFromFile("vendor/day5input-test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	seeds := []int{}
	min_location := 0
	seed_to_soil := [][]int{}
	soil_to_fertilizer := [][]int{}
	fertilizer_to_water := [][]int{}
	water_to_light := [][]int{}
	light_to_temperature := [][]int{}
	temperature_to_humidity := [][]int{}
	humidity_to_location := [][]int{}
	currently_loading := ""

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		if line == "" {
			continue
		}

		skip := false
		if strings.Contains(line, "seeds:") {
			seeds = collect_seeds(line)
			skip = true
		} else {
			currently_loading, skip = where_are_we(line, currently_loading)
		}

		if skip {
			fmt.Println("skipping line:", line)
			continue
		}

		// otherwise, parse the line and add to the appropriate map
		var destination_start int
		var source_start int
		var map_range int
		fmt.Sscanf(line, "%d %d %d", &destination_start, &source_start, &map_range)
		if destination_start == 0 && source_start == 0 && map_range == 0 {
			continue
		}
		range_map := []int{destination_start, source_start, map_range}

		switch currently_loading {
		case "seed_to_soil":
			seed_to_soil = append(seed_to_soil, range_map)
		case "soil_to_fertilizer":
			soil_to_fertilizer = append(soil_to_fertilizer, range_map)
		case "fertilizer_to_water":
			fertilizer_to_water = append(fertilizer_to_water, range_map)
		case "water_to_light":
			water_to_light = append(water_to_light, range_map)
		case "light_to_temperature":
			light_to_temperature = append(light_to_temperature, range_map)
		case "temperature_to_humidity":
			temperature_to_humidity = append(temperature_to_humidity, range_map)
		case "humidity_to_location":
			humidity_to_location = append(humidity_to_location, range_map)
		}
	}

	var soil int
	var fertilizer int
	var water int
	var light int
	var temperature int
	var humidity int
	var location int
	// now we have all the maps, so we can calculate the final location
	for index := 0; index < len(seeds)-1; index += 2 {
		seed_range := []int{seeds[index], seeds[index+1]}
		fmt.Println("seed_range:", seed_range)
		// first, calculate the soil
		for seed := seed_range[0]; seed < seed_range[0]+seed_range[1]+1; seed++ {
			soil = lookup_in(seed_to_soil, seed)
			fertilizer = lookup_in(soil_to_fertilizer, soil)
			water = lookup_in(fertilizer_to_water, fertilizer)
			light = lookup_in(water_to_light, water)
			temperature = lookup_in(light_to_temperature, light)
			humidity = lookup_in(temperature_to_humidity, temperature)
			location = lookup_in(humidity_to_location, humidity)

			if min_location == 0 || location < min_location {
				min_location = location
			}
		}
	}

	fmt.Println("min_location:", min_location)
}

func lookup_in(lookup_value [][]int, initial_value int) int {
	final_value := 0
	for _, range_map := range lookup_value {
		destination_start := range_map[0]
		source_start := range_map[1]
		map_range := range_map[2]
		final_value = calc_lookup(initial_value, destination_start, source_start, map_range)
		if final_value != 0 {
			break
		}
	}
	if final_value == 0 {
		final_value = initial_value
	}
	return final_value
}

func calc_lookup(seed int, destination_start int, source_start int, map_range int) int {
	// check if seed is within source_start..source_start+map_range
	if seed >= source_start && seed <= source_start+map_range {
		// if so, then calculate the destination
		seed_offset := seed - source_start
		destination := destination_start + seed_offset
		// fmt.Println("seed", seed, "is within range", source_start, "to", source_start+map_range, "so destination is", destination)
		return destination
	}
	return 0
}

func with_default(m map[int]int, key int) int {
	value := m[key]
	if value == 0 {
		value = key
	}
	return value
}

func collect_seeds(line string) []int {
	seeds := []int{}
	split := strings.Split(line, ":")
	if len(split) != 2 {
		fmt.Println("Error parsing line:", line)
		return seeds
	}
	seeds_to_find := strings.Split(split[1], " ")
	fmt.Println("seeds_to_find:", seeds_to_find)
	for _, seed := range seeds_to_find {
		if seed == "" {
			continue
		}
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("Error converting seed to integer:", err)
			continue
		}
		seeds = append(seeds, seedInt)
	}
	return seeds
}

func where_are_we(line string, currently_loading string) (string, bool) {
	now_loading := currently_loading
	if strings.Contains(line, "seeds:") {
		now_loading = "seeds"
	}

	if strings.Contains(line, "seed-to-soil map:") {
		now_loading = "seed_to_soil"
	}

	if strings.Contains(line, "soil-to-fertilizer map:") {
		now_loading = "soil_to_fertilizer"
	}

	if strings.Contains(line, "fertilizer-to-water map:") {
		now_loading = "fertilizer_to_water"
	}

	if strings.Contains(line, "water-to-light map:") {
		now_loading = "water_to_light"
	}

	if strings.Contains(line, "light-to-temperature map:") {
		now_loading = "light_to_temperature"
	}

	if strings.Contains(line, "temperature-to-humidity map:") {
		now_loading = "temperature_to_humidity"
	}

	if strings.Contains(line, "humidity-to-location map:") {
		now_loading = "humidity_to_location"
	}
	skip := now_loading != currently_loading
	return now_loading, skip
}
