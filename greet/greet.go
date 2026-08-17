// Package greet provides a friendly greeting.
package greet

// Hello returns a greeting for the given name.
func Hello(name string) string {
	if name == "" {
		return "Hello!"
	}
	return "Hello, " + name + "!"
}

// Farewell returns a farewell for the given name.
func Farewell(name string) string {
	if name == "" {
		return "Goodbye!"
	}
	return "Goodbye, " + name + "!"
}
