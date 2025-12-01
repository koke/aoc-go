package main

import (
	day1_2024 "aoc/2024/day1"
	day1_2025 "aoc/2025/day1"
	"aoc/utils"
	"fmt"
	"os"
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

	utils.DebugEnabled = false

	switch year {
	case "2024":
		if day == "1" {
			utils.RunDay(year, day, &day1_2024.Day{}, inputFile)
		}
	case "2025":
		if day == "1" {
			utils.RunDay(year, day, &day1_2025.Day{}, inputFile)
		}
	default:
		fmt.Printf("%s/%s not implemented yet\n", year, day)
		os.Exit(1)
	}
}
