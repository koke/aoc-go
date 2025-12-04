package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	fmt.Println("Running {{YEAR}} Day {{DAY}}")

	inputPath := utils.InitConfig().InputPath
	utils.RunPart("Part 1", inputPath, part1)
	utils.RunPart("Part 2", inputPath, part2)
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}
