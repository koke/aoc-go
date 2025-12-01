package day2_2024

import (
	"aoc/utils"
	"slices"
)

type Day struct{}

func isSafe(levels []int) bool {
	direction := utils.Sign(levels[1] - levels[0])
	for i := range len(levels) - 1 {
		x := levels[i]
		y := levels[i+1]
		diff := y - x
		if utils.Sign(diff) != direction {
			return false
		}
		diff = utils.Abs(diff)
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func (Day) Part1(input string) int {
	reports, err := utils.ReadMatrix(input)
	if err != nil {
		panic(err)
	}

	safe := 0
	for _, report := range reports {
		if isSafe(report) {
			safe++
		}
	}
	return safe
}

func (Day) Part2(input string) int {
	reports, err := utils.ReadMatrix(input)
	if err != nil {
		panic(err)
	}

	safe := 0
	for _, report := range reports {
		if isSafe(report) {
			safe++
		} else {
			for i := range len(report) {
				dampened := slices.Clone(report)
				dampened = slices.Delete(dampened, i, i+1)
				if isSafe(dampened) {
					safe++
					break
				}
			}
		}
	}
	return safe
}
