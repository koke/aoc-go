package day1

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines, err := utils.ReadLines(input)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return 0
	}

	var left, right []int
	for _, line := range lines {
		parts := strings.Fields(line)

		l, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error converting left value: %v\n", err)
			return 0
		}

		r, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error converting right value: %v\n", err)
			return 0
		}

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	diffSum := 0
	for i := range left {
		l := left[i]
		r := right[i]

		diff := l - r
		if diff < 0 {
			diff = -diff
		}

		utils.Debug("Pair %d: left=%d, right=%d, diff=%d", i+1, l, r, diff)
		diffSum += diff
	}

	return diffSum
}

func Part2(input string) int {
	return 0
}
