package utils

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type Config struct {
	Debug     bool
	TestMode  bool
	InputPath string
}

var config Config = Config{}

func InitConfig() Config {
	flag.BoolVar(&config.Debug, "debug", false, "Enable debug mode")
	flag.BoolVar(&config.TestMode, "test", false, "Use test input file")
	flag.Parse()

	config.InputPath = getInputPath()

	return config
}

func RunPart[T any](label string, input T, fn func(T) int) {
	start := time.Now()
	result := fn(input)
	elapsed := time.Since(start)
	fmt.Printf("%s: %d (Time: %s)\n", label, result, elapsed)
}

func getInputPath() string {
	inputFile := "input.txt"
	if config.TestMode {
		inputFile = "test.txt"
	}

	// Get the caller's source file location to determine the directory
	_, callerFile, _, ok := runtime.Caller(2)
	if !ok {
		return inputFile
	}

	// Get the directory containing the caller's source file
	callerDir := filepath.Dir(callerFile)
	return filepath.Join(callerDir, inputFile)
}
