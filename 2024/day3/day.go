package day3_2024

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Day struct{}

func (d Day) Run(input string) {
	utils.RunPart("Part 1", input, part1)
	utils.RunPart("Part 2", input, part2)
}

func part1(input string) int {
	memoryBytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	memory := string(memoryBytes)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(memory, -1)

	result := 0
	for _, match := range matches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		result += x * y
	}
	return result
}

func part2(input string) int {
	memoryBytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	memory := string(memoryBytes)

	re := regexp.MustCompile(`(mul)\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(memory, -1)
	fmt.Println(matches)

	result := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if match[1] == "mul" && enabled {
			x, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(match[3])
			if err != nil {
				panic(err)
			}
			result += x * y
		}
	}
	return result
}
