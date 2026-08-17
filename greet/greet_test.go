package greet

import "testing"

func TestHello(t *testing.T) {
	got := Hello("World")
	want := "Hello, World!"
	if got != want {
		t.Errorf("Hello(%q) = %q, want %q", "World", got, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	got := Hello("")
	want := "Hello, !"
	if got != want {
		t.Errorf("Hello(%q) = %q, want %q", "", got, want)
	}
}
