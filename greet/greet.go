// Package greet provides a friendly greeting.
package greet

import "strings"

// normalizeName trims surrounding whitespace from name.
func normalizeName(name string) string {
	return strings.TrimSpace(name)
}

// Hello returns a greeting for the given name.
func Hello(name string) string {
	name = normalizeName(name)
	if name == "" {
		return "Hello!"
	}
	return "Hello, " + name + "!"
}

// Farewell returns a farewell for the given name.
func Farewell(name string) string {
	name = normalizeName(name)
	if name == "" {
		return "Goodbye!"
	}
	return "Goodbye, " + name + "!"
}
