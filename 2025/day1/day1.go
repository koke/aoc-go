package day1

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

type Day struct{}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Returns x modulo 100, wrapping around correctly for negative numbers
func circlemod(x int) int {
	x = x % 100
	if x < 0 {
		x += 100
	}
	return x
}

func parseInput(input string) ([]int, error) {
	lines, err := utils.ReadLines(input)
	if err != nil {
		return nil, err
	}

	var numbers []int
	for _, line := range lines {
		direction := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		if direction == 'L' {
			count = -count
		}
		numbers = append(numbers, count)
	}

	return numbers, nil
}

func turnDial(current int, rotation int) (int, int) {
	// Basically, we want to add current+coration modulo 100
	// but we also want to count how many times we pass 0, not including the start point
	final := current + rotation
	clicks := 0

	// If we went from nevative to positive or positive to negative, we passed 0 at least once
	if current < 0 && final >= 0 || current > 0 && final <= 0 {
		clicks++
	}

	// Then, if we moved beyond the 100 range, we count how many times we passed 0
	clicks += abs(final / 100)

	return circlemod(final), clicks
}

func (Day) Part1(input string) int {
	rotations, err := parseInput(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return 0
	}

	dial := 50
	password := 0

	for _, count := range rotations {
		oldDial := dial
		dial = (dial + count) % 100
		if dial < 0 {
			dial += 100
		}
		if dial == 0 {
			password += 1
		}
		utils.Debug("%d -- %d --> %d", oldDial, count, dial)
	}

	return password
}

func (Day) Part2(input string) int {
	rotations, err := parseInput(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return 0
	}

	dial := 50
	password := 0

	for _, count := range rotations {
		oldDial := dial
		var clicks int
		dial, clicks = turnDial(dial, count)
		click := ""
		password += clicks
		for range clicks {
			click += "click "
		}
		utils.Debug("%d -- %d --> %d %s", oldDial, count, dial, click)
	}

	return password
}
