package day1

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func Part1(input string) int {
	lines, err := utils.ReadLines(input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return 0
	}

	dial := 50
	password := 0

	for _, line := range lines {
		direction := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error converting count:", err)
		}
		if direction == 'L' {
			count = -count
		}

		oldDial := dial
		dial = (dial + count) % 100
		if dial < 0 {
			dial += 100
		}
		if dial == 0 {
			password += 1
		}
		utils.Debug("%d -- %s --> %d", oldDial, line, dial)
	}

	return password
}

func Part2(input string) int {
	return 0
}
