package day2_2025

import (
	"fmt"
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
		// fmt.Printf("Checking number %d: left='%s', right='%s', valid=%v\n", n, left, right, valid)
		if !valid {
			return false
		}
	} else {
		// fmt.Printf("Number %d has odd length, automatically valid\n", n)
	}
	return true
}

func (Day) Part1(input string) int {
	ranges := parseInput(input)
	fmt.Printf("Parsed ranges: %+v\n", ranges)

	invalidSum := 0

	for _, r := range ranges {
		fmt.Printf("Processing range: %d-%d\n", r.low, r.high)
		for n := r.low; n <= r.high; n++ {
			if isValid(n) {
				// fmt.Printf("Valid number: %d\n", n)
			} else {
				fmt.Printf("Invalid number: %d\n", n)
				invalidSum += n
			}
		}
	}

	return invalidSum
}

func (Day) Part2(input string) int {
	return 0
}
