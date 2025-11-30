package main

import (
	"aoc/2024/day1"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <year> <day> [test]")
		fmt.Println("Example: go run . 2024 1")
		fmt.Println("Example: go run . 2024 1 test")
		os.Exit(1)
	}

	year := os.Args[1]
	day := os.Args[2]

	// Determine if using test input
	inputFile := "input.txt"
	if len(os.Args) > 3 && os.Args[3] == "test" {
		inputFile = "test.txt"
	}

	inputPath := filepath.Join(year, fmt.Sprintf("day%s", day), inputFile)

	if year == "2024" {
		if day == "1" {
			fmt.Printf("Running 2024 Day 1\n")
			fmt.Println("Part 1:", day1.Part1(inputPath))
			// fmt.Println("Part 2:", day1.Part2(inputPath))
		}
	} else {
		fmt.Printf("%s/%s not implemented yet\n", year, day)
		os.Exit(1)
	}
}
