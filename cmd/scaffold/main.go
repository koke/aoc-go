package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func findLatestDay() (string, string) {
	entries, err := os.ReadDir(".")
	if err != nil {
		return "2024", "1"
	}

	var years []string
	for _, entry := range entries {
		if entry.IsDir() && len(entry.Name()) == 4 {
			if _, err := strconv.Atoi(entry.Name()); err == nil {
				years = append(years, entry.Name())
			}
		}
	}

	if len(years) == 0 {
		return "2024", "1"
	}

	sort.Strings(years)
	latestYear := years[len(years)-1]

	// Find latest day in that year
	yearEntries, err := os.ReadDir(latestYear)
	if err != nil {
		return latestYear, "1"
	}

	var days []int
	for _, entry := range yearEntries {
		if entry.IsDir() {
			if day, err := strconv.Atoi(entry.Name()); err == nil {
				days = append(days, day)
			}
		}
	}

	if len(days) == 0 {
		return latestYear, "1"
	}

	sort.Ints(days)
	nextDay := days[len(days)-1] + 1

	return latestYear, strconv.Itoa(nextDay)
}

func main() {
	// Find latest year/day
	suggestedYear, suggestedDay := findLatestDay()
	suggestion := fmt.Sprintf("%s/%s", suggestedYear, suggestedDay)

	// Prompt user
	fmt.Printf("Create new day [%s]: ", suggestion)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Use suggestion if empty
	if input == "" {
		input = suggestion
	}

	// Parse input
	parts := strings.Split(input, "/")
	if len(parts) != 2 {
		fmt.Println("Invalid format. Expected: year/day (e.g., 2025/4)")
		os.Exit(1)
	}

	year := parts[0]
	day := parts[1]

	// Create directory
	dir := filepath.Join(year, day)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Copy and process template files
	templateDir := "template"

	// Read template day.go
	templatePath := filepath.Join(templateDir, "day.go")
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		os.Exit(1)
	}

	// Replace placeholders
	content := string(templateContent)
	content = strings.ReplaceAll(content, "{{YEAR}}", year)
	content = strings.ReplaceAll(content, "{{DAY}}", day)

	// Write day.go
	dayPath := filepath.Join(dir, "day.go")
	if err := os.WriteFile(dayPath, []byte(content), 0644); err != nil {
		fmt.Printf("Error writing day.go: %v\n", err)
		os.Exit(1)
	}

	// Copy empty input files
	os.WriteFile(filepath.Join(dir, "input.txt"), []byte(""), 0644)
	os.WriteFile(filepath.Join(dir, "test.txt"), []byte(""), 0644)

	fmt.Printf("✓ Created %s/%s/\n", year, day)
	fmt.Printf("Run with: go run %s/%s\n", year, day)
	fmt.Printf("Run test: go run %s/%s test\n", year, day)
}
