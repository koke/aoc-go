package main

import (
	day1_2024 "aoc/2024/day1"
	day2_2024 "aoc/2024/day2"
	day1_2025 "aoc/2025/day1"
	"aoc/utils"
	"fmt"
	"os"
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

	var day utils.Day

	switch input {
	case "2024/1":
		day = &day1_2024.Day{}
	case "2024/2":
		day = &day2_2024.Day{}
	case "2025/1":
		day = &day1_2025.Day{}
	default:
		fmt.Printf("%s not implemented yet\n", input)
		os.Exit(1)
	}

	utils.RunDay(input, day, inputFile)
}
