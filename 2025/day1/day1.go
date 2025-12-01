package day1_2025

import (
	"aoc/utils"
	"strconv"
)

type Day struct{}

// Returns x modulo 100, wrapping around correctly for negative numbers
func circlemod(x int) int {
	x = x % 100
	if x < 0 {
		x += 100
	}
	return x
}

func parseInput(input string) []int {
	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	var numbers []int
	for _, line := range lines {
		direction := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if direction == 'L' {
			count = -count
		}
		numbers = append(numbers, count)
	}

	return numbers
}

func turnDial(current int, rotation int) (int, int) {
	// Basically, we want to add current+rotation modulo 100
	// but we also want to count how many times we pass 0, not including the start point
	final := current + rotation
	clicks := 0

	// If we went from negative to positive or positive to negative, we passed 0 at least once
	if current < 0 && final >= 0 || current > 0 && final <= 0 {
		clicks++
	}

	// Then, if we moved beyond the 100 range, we count how many times we passed 0
	clicks += utils.Abs(final / 100)

	return circlemod(final), clicks
}

func (Day) Part1(input string) int {
	rotations := parseInput(input)

	dial := 50
	password := 0

	for _, count := range rotations {
		dial, _ = turnDial(dial, count)

		if dial == 0 {
			password += 1
		}
	}

	return password
}

func (Day) Part2(input string) int {
	rotations := parseInput(input)

	dial := 50
	password := 0

	for _, count := range rotations {
		var clicks int
		dial, clicks = turnDial(dial, count)
		password += clicks
	}

	return password
}
