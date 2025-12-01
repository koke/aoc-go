package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type Day interface {
	Part1(inputPath string) int
	Part2(inputPath string) int
}

func RunDay(input string, d Day, inputFile string) {
	parts := strings.Split(input, "/")
	year := parts[0]
	day := parts[1]

	inputPath := filepath.Join(year, fmt.Sprintf("day%s", day), inputFile)

	fmt.Printf("Running %s Day %s\n", year, day)

	start1 := time.Now()
	result1 := d.Part1(inputPath)
	elapsed1 := time.Since(start1)
	fmt.Printf("Part 1: %d (Time: %s)\n", result1, elapsed1)

	start2 := time.Now()
	result2 := d.Part2(inputPath)
	elapsed2 := time.Since(start2)
	fmt.Printf("Part 2: %d (Time: %s)\n", result2, elapsed2)
}
