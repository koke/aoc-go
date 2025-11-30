package utils

import "fmt"

var DebugEnabled = false

func Debug(format string, args ...interface{}) {
	if DebugEnabled {
		fmt.Printf("[DEBUG] "+format+"\n", args...)
	}
}
