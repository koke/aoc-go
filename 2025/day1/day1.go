package day1

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

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
	step := 1
	clicks := 0
	if rotation < 0 {
		step = -1
		rotation = -rotation
	}
	for range rotation {
		current += step
		if current == 100 {
			current = 0
		}
		if current == 0 {
			clicks++
		}
		if current == -1 {
			current = 99
		}
	}
	return current, clicks
}

func Part1(input string) int {
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

func Part2(input string) int {
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
