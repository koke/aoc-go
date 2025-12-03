package day2_2025

import (
	"os"
	"strconv"
	"strings"
)

type Day struct{}

type Range struct {
	low, high int
}

func parseInput(filename string) []Range {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var ranges []Range
	for _, rangeStr := range strings.Split(strings.TrimSpace(string(data)), ",") {
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

func isValid(n int) bool {
	number := strconv.Itoa(n)
	numberLen := len(number)

	if numberLen%2 == 0 {
		left := number[:numberLen/2]
		right := number[numberLen/2:]
		valid := left != right

		if !valid {
			return false
		}
	}
	return true
}

func isValid2(n int) bool {
	number := strconv.Itoa(n)
	numberLen := len(number)

	for chunkSize := range numberLen / 2 {
		repeatCount := numberLen / (chunkSize + 1)
		chunk := number[:chunkSize+1]
		repeated := strings.Repeat(chunk, repeatCount)
		valid := repeated != number
		if !valid {
			return false
		}
	}
	return true
}

func (Day) Part1(input string) int {
	ranges := parseInput(input)

	invalidSum := 0

	for _, r := range ranges {
		for n := r.low; n <= r.high; n++ {
			if isValid(n) {
				// fmt.Printf("Valid number: %d\n", n)
			} else {
				invalidSum += n
			}
		}
	}

	return invalidSum
}

func (Day) Part2(input string) int {
	ranges := parseInput(input)

	invalidSum := 0

	for _, r := range ranges {
		for n := r.low; n <= r.high; n++ {
			if !isValid2(n) {
				invalidSum += n
			}
		}
	}

	return invalidSum
}
