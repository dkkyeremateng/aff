// Package greet provides a friendly greeting.
package greet

import "strings"

// normalizeName trims surrounding whitespace from name and reports whether the
// trimmed result is empty.
func normalizeName(name string) (string, bool) {
	name = strings.TrimSpace(name)
	return name, name == ""
}

// Hello returns a greeting for the given name.
func Hello(name string) string {
	name, isEmpty := normalizeName(name)
	if isEmpty {
		return "Hello!"
	}
	return "Hello, " + name + "!"
}

// Farewell returns a farewell for the given name.
func Farewell(name string) string {
	name, isEmpty := normalizeName(name)
	if isEmpty {
		return "Goodbye!"
	}
	return "Goodbye, " + name + "!"
}
