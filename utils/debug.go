package utils

import "fmt"

func Debug(format string, args ...any) {
	if config.Debug {
		fmt.Printf("[DEBUG] "+format+"\n", args...)
	}
}
