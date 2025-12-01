package main

import (
	day1_2024 "aoc/2024/day1"
	day1_2025 "aoc/2025/day1"
	"aoc/utils"
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

	utils.DebugEnabled = true

	switch year {
	case "2024":
		if day == "1" {
			fmt.Printf("Running 2024 Day 1\n")
			fmt.Println("Part 1:", day1_2024.Part1(inputPath))
			fmt.Println("Part 2:", day1_2024.Part2(inputPath))
		}
	case "2025":
		if day == "1" {
			fmt.Println("Part 1:", day1_2025.Part1(inputPath))
			fmt.Println("Part 2:", day1_2025.Part2(inputPath))
		}
	default:
		fmt.Printf("%s/%s not implemented yet\n", year, day)
		os.Exit(1)
	}
}
