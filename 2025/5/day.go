package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strings"
)

type database struct {
	ranges      []utils.Range
	ingredients []int
}

func main() {
	fmt.Println("Running 2025 Day 5")

	inputPath := utils.InitConfig().InputPath
	db := parseInput(inputPath)
	utils.RunPart("Part 1", db, part1)
	utils.RunPart("Part 2", db, part2)
}

func part1(db database) int {
	freshCount := 0
	for _, ingredient := range db.ingredients {
		if isFresh(ingredient, db) {
			freshCount++
		}
	}
	return freshCount
}

func part2(db database) int {
	// The naive approach won't work as the input has very large ranges.
	// We need to find all the overlapping ranges and combine them.

	sort.Slice(db.ranges, func(i, j int) bool {
		return db.ranges[i].Low < db.ranges[j].Low
	})
	combinedRanges := []utils.Range{db.ranges[0]}
	utils.Debug("Combining ranges: %v", db.ranges)

	for i := 1; i < len(db.ranges); i++ {
		r := db.ranges[i]
		last := &combinedRanges[len(combinedRanges)-1]

		if last.Overlaps(r) {
			utils.Debug("Combining overlapping ranges %d-%d and %d-%d", last.Low, last.High, r.Low, r.High)
			last.High = max(last.High, r.High)
		} else {
			utils.Debug("Adding non-overlapping range %d-%d", r.Low, r.High)
			combinedRanges = append(combinedRanges, r)
		}
		utils.Debug("Combined ranges so far: %v", combinedRanges)
	}

	totalFresh := 0
	for _, r := range combinedRanges {
		totalFresh += r.High - r.Low + 1
	}

	return totalFresh
}

func parseInput(input string) database {
	step := 0 // 0: ranges, 1: ingredients
	lines := utils.ReadLines(input)

	db := database{}

	for _, line := range lines {
		if line == "" {
			if step > 0 {
				panic("unexpected blank line")
			}
			step++
		} else if step == 0 {
			// Parse ranges
			parts := strings.Split(line, "-")
			low := utils.Atoi(parts[0])
			high := utils.Atoi(parts[1])
			db.ranges = append(db.ranges, utils.Range{Low: low, High: high})
		} else {
			// Parse ingredients
			ingredient := utils.Atoi(line)
			db.ingredients = append(db.ingredients, ingredient)
		}
	}

	return db
}

func isFresh(ingredient int, db database) bool {
	for _, r := range db.ranges {
		if r.Contains(ingredient) {
			utils.Debug("Ingredient %d is fresh in range %d-%d", ingredient, r.Low, r.High)
			return true
		}
	}
	utils.Debug("Ingredient %d is spoiled", ingredient)
	return false
}
