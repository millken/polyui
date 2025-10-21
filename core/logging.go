package core

import "fmt"

// DebugEnabled controls whether debug output is printed.
// Set to true in tests or main when you want verbose output.
var DebugEnabled = false

// Debugf prints a formatted debug message when DebugEnabled is true.
func Debugf(format string, args ...interface{}) {
	if DebugEnabled {
		fmt.Printf(format, args...)
	}
}

// Debugln prints a newline-terminated debug message when DebugEnabled is true.
func Debugln(args ...interface{}) {
	if DebugEnabled {
		fmt.Println(args...)
	}
}
