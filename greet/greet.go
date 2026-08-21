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
		return "Hello there!"
	}
	return "Hello, " + name + "!"
}

// Farewell returns a farewell for the given name.
func Farewell(name string) string {
	name = normalizeName(name)
	if name == "" {
		return "Goodbye there!"
	}
	return "Goodbye, " + name + "!"
}

// Welcome returns a welcome for the given name.
func Welcome(name string) string {
	name = normalizeName(name)
	if name == "" {
		return "Welcome!"
	}
	return "Welcome, " + name + "!"
}
