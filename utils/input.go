package utils

import (
	"os"
	"strings"
)

// ReadLines reads a file and returns lines as a slice of strings
func ReadLines(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines
}

func ReadMatrix(filename string) ([][]int, error) {
	lines := ReadLines(filename)

	matrix := make([][]int, len(lines))
	for i, line := range lines {
		values := strings.Fields(line)
		row := make([]int, len(values))
		for j, v := range values {
			row[j] = Atoi(v)
		}
		matrix[i] = row
	}
	return matrix, nil
}
