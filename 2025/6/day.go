package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Running 2025 Day 6")

	inputPath := utils.InitConfig().InputPath
	utils.RunPart("Part 1", inputPath, part1)
	utils.RunPart("Part 2", inputPath, part2)
}

type operation string

const (
	addition       operation = "+"
	multiplication operation = "*"
)

type problem struct {
	values   []int
	operator operation
}

func parseInput(input string) []problem {
	lines := utils.ReadLines(input)
	problems := make([]problem, len(strings.Fields(lines[0])))
	for _, line := range lines {
		values := strings.Fields(line)
		for i, v := range values {
			number, err := strconv.Atoi(v)
			if err != nil {
				problems[i].operator = operation(v)
			} else {
				problems[i].values = append(problems[i].values, number)
			}
		}
	}
	return problems
}

func solve(prob problem) int {
	solution := prob.values[0]
	for _, v := range prob.values[1:] {
		switch prob.operator {
		case addition:
			solution += v
		case multiplication:
			solution *= v
		default:
			panic("unknown operator")
		}
	}
	return solution
}

func part1(input string) int {
	problems := parseInput(input)

	solution := 0
	for _, prob := range problems {
		solution += solve(prob)
	}
	return solution
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	lastLine := lines[len(lines)-1]

	solution := 0
	values := []int{}
	// values := make([]int, len(lines))
	for col := len(lines[0]) - 1; col >= 0; col-- {
		colValue := 0
		for row := range len(lines) - 1 {
			value, err := strconv.Atoi(string(lines[row][col]))
			if err == nil {
				colValue = colValue*10 + value
			}
		}
		utils.Debug("Column %d value: %d", col, colValue)
		if colValue > 0 {
			values = append(values, colValue)
		}
		var op byte
		if len(lastLine) > col {
			op = lastLine[col]
			if op != '+' && op != '*' {
				continue
			}
			problem := problem{
				values:   values,
				operator: operation(op),
			}
			result := solve(problem)
			utils.Debug("Values: %v", values)
			utils.Debug("Operator at column %d: %c. Result: %d", col, op, result)
			solution += result
			values = []int{}
		}
	}
	return solution
}
