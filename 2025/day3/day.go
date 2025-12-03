package day3_2025

import (
	"aoc/utils"
	"fmt"
)

type Day struct{}

func maxJoltage(bank string) int {
	battery1 := '0'
	battery2 := '0'

	for pos, battery := range bank {
		if battery > battery1 && pos+1 < len(bank) {
			battery1, battery2 = battery, 0
		} else if battery > battery2 {
			battery2 = battery
		}
	}
	return int(battery1-'0')*10 + int(battery2-'0')
}

func (Day) Part1(input string) int {
	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	totalJoltage := 0
	for _, line := range lines {
		joltage := maxJoltage(line)
		fmt.Println("Bank:", line, "Max Joltage:", joltage)
		totalJoltage += joltage
	}
	return totalJoltage
}

func (Day) Part2(input string) int {
	return 0
}
