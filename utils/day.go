package utils

import (
	"fmt"
	"time"
)

type Day interface {
	Run(inputPath string)
}

func RunPart[T any](label string, input T, fn func(T) int) {
	start := time.Now()
	result := fn(input)
	elapsed := time.Since(start)
	fmt.Printf("%s: %d (Time: %s)\n", label, result, elapsed)
}
