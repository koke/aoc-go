package utils

import (
	"os"
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
