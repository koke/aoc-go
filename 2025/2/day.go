package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Running 2025 Day 2")

	inputPath := utils.InitConfig().InputPath
	ranges := parseInput(inputPath)

	utils.RunPart("Part 1", ranges, part1)
	utils.RunPart("Part 2", ranges, part2)
}

func parseInput(filename string) []utils.Range {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var ranges []utils.Range
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

		ranges = append(ranges, utils.Range{Low: lower, High: higher})
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

func validateInput(ranges []utils.Range, maxRepeat int) int {
	invalidSum := 0

	for _, r := range ranges {
		for n := r.Low; n <= r.High; n++ {
			if !isValid(n, maxRepeat) {
				invalidSum += n
			}
		}
	}

	return invalidSum
}

func part1(ranges []utils.Range) int {
	return validateInput(ranges, 2)
}

func part2(ranges []utils.Range) int {
	return validateInput(ranges, math.MaxInt)
}
