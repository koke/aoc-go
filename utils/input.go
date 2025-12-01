package utils

import (
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns lines as a slice of strings
func ReadLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines, nil
}

func ReadMatrix(filename string) ([][]int, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	matrix := make([][]int, len(lines))
	for i, line := range lines {
		values := strings.Fields(line)
		row := make([]int, len(values))
		for j, v := range values {
			row[j], err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		}
		matrix[i] = row
	}
	return matrix, nil
}
