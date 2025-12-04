package day2_2025

import (
	"aoc/utils"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day struct{}

func (d Day) Run(input string) {
	utils.RunPart("Part 1", input, part1)
	utils.RunPart("Part 2", input, part2)
}

type Range struct {
	low, high int
}

func parseInput(filename string) []Range {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var ranges []Range
	for rangeStr := range strings.SplitSeq(strings.TrimSpace(string(data)), ",") {
		rangeParts := strings.Split(rangeStr, "-")
		lower, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			panic(err)
		}
		higher, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, Range{low: lower, high: higher})
	}
	return ranges
}

func isValid(n int, maxRepeat int) bool {
	number := strconv.Itoa(n)
	numberLen := len(number)

	for chunkSize := numberLen / 2; chunkSize >= 1; chunkSize-- {
		repeatCount := numberLen / (chunkSize)
		if repeatCount > maxRepeat {
			return true
		}
		chunk := number[:chunkSize]
		repeated := strings.Repeat(chunk, repeatCount)
		valid := repeated != number
		if !valid {
			return false
		}
	}
	return true
}

func validateInput(input string, maxRepeat int) int {
	ranges := parseInput(input)

	invalidSum := 0

	for _, r := range ranges {
		for n := r.low; n <= r.high; n++ {
			if !isValid(n, maxRepeat) {
				invalidSum += n
			}
		}
	}

	return invalidSum
}

func part1(input string) int {
	return validateInput(input, 2)
}

func part2(input string) int {
	return validateInput(input, math.MaxInt)
}
