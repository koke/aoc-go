package main

import (
	"aoc/utils"
	"fmt"
)

var config utils.Config

func main() {
	fmt.Println("Running 2025 Day 4")

	config = utils.InitConfig()
	inputPath := config.InputPath
	lines := utils.ReadLines(inputPath)
	utils.RunPart("Part 1", lines, part1)
	utils.RunPart("Part 2", lines, part2)
	utils.RunPart("Part 2 (Optimized)", lines, part2Optimized)
}

func parseGrid(input []string) [][]bool {
	grid := make([][]bool, len(input))
	for lineIndex, line := range input {
		for _, char := range line {
			grid[lineIndex] = append(grid[lineIndex], char == '@')
		}
	}
	return grid
}

func isAccessible(grid [][]bool, x, y int) bool {
	if !grid[y][x] {
		return false
	}

	rollCount := 0
	for nx := x - 1; nx <= x+1; nx++ {
		for ny := y - 1; ny <= y+1; ny++ {
			if nx >= 0 && ny >= 0 && ny < len(grid) && nx < len(grid[0]) && (nx != x || ny != y) {
				hasRoll := grid[ny][nx]
				if hasRoll {
					rollCount++
				}
				utils.Debug("Checking %d,%d for %d,%d: %v rollCount: %d", nx, ny, x, y, hasRoll, rollCount)
			}
		}
	}
	return rollCount < 4
}

func calculateAccessibilityGrid(grid [][]bool) ([][]bool, int) {
	rows := len(grid)
	cols := len(grid[0])

	accessibilityGrid := make([][]bool, rows)
	totalAccessible := 0
	for y := range rows {
		accessibilityGrid[y] = make([]bool, cols)
		for x := range cols {
			accessible := isAccessible(grid, x, y)
			accessibilityGrid[y][x] = accessible
			if accessible {
				totalAccessible++
			}
			utils.Debug("Cell %d,%d accessible: %v", x, y, accessible)
		}
	}
	return accessibilityGrid, totalAccessible
}

func calculateAccessibilityGridOptimized(grid [][]bool) ([][]bool, int) {
	rows := len(grid)
	if rows == 0 {
		return [][]bool{}, 0
	}
	cols := len(grid[0])

	accessibilityGrid := make([][]bool, rows)
	totalAccessible := 0

	for i := range rows {
		accessibilityGrid[i] = make([]bool, cols)
		for j := range cols {
			if !grid[i][j] {
				continue // Cell itself is not occupied
			}

			// Count neighbors
			count := 0
			if i > 0 {
				if j > 0 && grid[i-1][j-1] {
					count++
				}
				if grid[i-1][j] {
					count++
				}
				if j < cols-1 && grid[i-1][j+1] {
					count++
				}
			}
			if j > 0 && grid[i][j-1] {
				count++
			}
			if j < cols-1 && grid[i][j+1] {
				count++
			}
			if i < rows-1 {
				if j > 0 && grid[i+1][j-1] {
					count++
				}
				if grid[i+1][j] {
					count++
				}
				if j < cols-1 && grid[i+1][j+1] {
					count++
				}
			}

			if count < 4 {
				accessibilityGrid[i][j] = true
				totalAccessible++
			}
		}
	}

	return accessibilityGrid, totalAccessible
}

func printGrid(grid [][]bool, accessibilityGrid [][]bool) {
	if !config.Debug {
		return
	}
	for y := range grid {
		for x := range grid[0] {
			char := '.'
			if grid[y][x] {
				char = '@'
			}
			if y < len(accessibilityGrid) && x < len(accessibilityGrid[0]) && accessibilityGrid[y][x] {
				char = 'x'
			}
			fmt.Printf("%c", char)
		}
		fmt.Printf("\n")
	}
}

func removeAccessibleRolls(grid [][]bool, accessible [][]bool) [][]bool {
	rows := len(grid)
	cols := len(grid[0])

	newGrid := make([][]bool, rows)
	for y := range rows {
		newGrid[y] = make([]bool, cols)
		for x := range cols {
			hasRoll := grid[y][x] && !accessible[y][x]
			newGrid[y][x] = hasRoll
		}
	}
	return newGrid
}

func part1(input []string) int {
	grid := parseGrid(input)
	printGrid(grid, [][]bool{})
	accessibilityGrid, result := calculateAccessibilityGrid(grid)
	printGrid(grid, accessibilityGrid)
	return result
}

func part2(input []string) int {
	grid := parseGrid(input)
	removed := 0
	for i := 0; ; i++ {
		accessibilityGrid, result := calculateAccessibilityGrid(grid)
		utils.Debug("Iteration %d: removed %d rolls", i, result)
		printGrid(grid, accessibilityGrid)
		removed += result
		if result == 0 {
			return removed
		}
		grid = removeAccessibleRolls(grid, accessibilityGrid)
	}
}

func part2Optimized(input []string) int {
	grid := parseGrid(input)
	removed := 0
	for i := 0; ; i++ {
		accessibilityGrid, result := calculateAccessibilityGridOptimized(grid)
		utils.Debug("Iteration %d: removed %d rolls", i, result)
		printGrid(grid, accessibilityGrid)
		removed += result
		if result == 0 {
			return removed
		}
		grid = removeAccessibleRolls(grid, accessibilityGrid)
	}
}
