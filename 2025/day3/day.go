package day3_2025

import (
	"aoc/utils"
	"strings"
)

type Day struct{}

type partInput struct {
	lines       []string
	batteriesOn int
}

func (d Day) Run(input string) {
	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	utils.RunPart("Part 1", partInput{lines, 2}, part)
	utils.RunPart("Part 2", partInput{lines, 12}, part)
}

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

func part(input partInput) int {
	totalJoltage := 0
	for _, line := range input.lines {
		joltage := maxJoltage(line, input.batteriesOn)
		// fmt.Println("Bank:", line, "Max Joltage:", joltage)
		totalJoltage += joltage
	}
	return totalJoltage
}
