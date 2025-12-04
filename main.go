package main

import (
	day1_2024 "aoc/2024/1"
	day2_2024 "aoc/2024/2"
	day3_2024 "aoc/2024/3"
	day1_2025 "aoc/2025/1"
	day2_2025 "aoc/2025/2"
	day3_2025 "aoc/2025/3"
	"aoc/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <year> <day> [test]")
		fmt.Println("Example: go run . 2024/1")
		fmt.Println("Example: go run . 2024/1 test")
		os.Exit(1)
	}

	input := os.Args[1]

	// Determine if using test input
	inputFile := "input.txt"
	if len(os.Args) > 2 && os.Args[2] == "test" {
		inputFile = "test.txt"
	}

	utils.DebugEnabled = false

	// Parse year/day and build input path
	parts := strings.Split(input, "/")
	year := parts[0]
	dayNum := parts[1]
	inputPath := filepath.Join(year, dayNum, inputFile)

	fmt.Printf("Running %s Day %s\n", year, dayNum)

	var day utils.Day

	switch input {
	case "2024/1":
		day = &day1_2024.Day{}
	case "2024/2":
		day = &day2_2024.Day{}
	case "2024/3":
		day = &day3_2024.Day{}
	case "2025/1":
		day = &day1_2025.Day{}
	case "2025/2":
		day = &day2_2025.Day{}
	case "2025/3":
		day = &day3_2025.Day{}
	default:
		fmt.Printf("%s not implemented yet\n", input)
		os.Exit(1)
	}

	day.Run(inputPath)
}
