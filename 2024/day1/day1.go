package day1_2024

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day struct{}

func readPairs(input string) (left []int, right []int, err error) {
	lines, err := utils.ReadLines(input)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading input: %v", err)
	}

	for _, line := range lines {
		parts := strings.Fields(line)

		l, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting left value: %v", err)
		}

		r, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting right value: %v", err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right, nil
}

func (Day) Part1(input string) int {
	left, right, err := readPairs(input)
	if err != nil {
		panic(err)
	}

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

func (Day) Part2(input string) int {
	left, right, err := readPairs(input)
	if err != nil {
		panic(err)
	}

	similarityScore := 0
	for i, l := range left {
		matches := 0
		for _, r := range right {
			if l == r {
				matches++
			}
		}
		utils.Debug("[%d] %d matches %d times", i, l, matches)
		similarityScore += matches * l
	}

	return similarityScore
}
