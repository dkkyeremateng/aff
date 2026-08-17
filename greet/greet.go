// Package greet provides a friendly greeting.
package greet

import "strings"

// Hello returns a greeting for the given name.
func Hello(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Hello there!"
	}
	return "Hello, " + name + "!"
}

// Farewell returns a farewell for the given name.
func Farewell(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Goodbye there!"
	}
	return "Goodbye, " + name + "!"
}
