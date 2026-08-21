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
	want := "Hello there!"
	if got != want {
		t.Errorf("Hello(%q) = %q, want %q", "", got, want)
	}
}

func TestFarewell(t *testing.T) {
	got := Farewell("World")
	want := "Goodbye, World!"
	if got != want {
		t.Errorf("Farewell(%q) = %q, want %q", "World", got, want)
	}
}

func TestFarewellEmpty(t *testing.T) {
	got := Farewell("")
	want := "Goodbye there!"
	if got != want {
		t.Errorf("Farewell(%q) = %q, want %q", "", got, want)
	}
}

func TestWelcome(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "Ada", want: "Welcome, Ada!"},
		{name: " Ada ", want: "Welcome, Ada!"},
		{name: "", want: "Welcome!"},
	}

	for _, test := range tests {
		got := Welcome(test.name)
		if got != test.want {
			t.Errorf("Welcome(%q) = %q, want %q", test.name, got, test.want)
		}
	}
}
