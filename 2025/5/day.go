package main

import (
	"aoc/utils"
	"fmt"
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
	return 0
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
