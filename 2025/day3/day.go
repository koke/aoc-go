package day3_2025

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Day struct{}

func maxJoltage(bank string, batteriesOn int) int {
	joltage := []byte(strings.Repeat("0", batteriesOn))

	for bankPos := range len(bank) {
		minPos := max(0, batteriesOn-(len(bank)-bankPos))
		for joltagePos := minPos; joltagePos < batteriesOn; joltagePos++ {
			if bank[bankPos] > joltage[joltagePos] {
				prevJoltage := make([]byte, len(joltage))
				copy(prevJoltage, joltage)
				joltage[joltagePos] = bank[bankPos]
				// Clear everything to the right of joltagePos
				for k := joltagePos + 1; k < len(joltage); k++ {
					joltage[k] = '0'
				}
				break
			}
		}
	}

	return utils.Atoi(string(joltage))
}

func (Day) Part1(input string) int {
	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	totalJoltage := 0
	for _, line := range lines {
		joltage := maxJoltage(line, 2)
		fmt.Println("Bank:", line, "Max Joltage:", joltage)
		totalJoltage += joltage
	}
	return totalJoltage
}

func (Day) Part2(input string) int {
	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	totalJoltage := 0
	for _, line := range lines {
		joltage := maxJoltage(line, 12)
		fmt.Println("Bank:", line, "Max Joltage:", joltage)
		totalJoltage += joltage
	}
	return totalJoltage
}
