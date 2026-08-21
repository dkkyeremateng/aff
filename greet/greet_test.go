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

func TestThanks(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "direct name", input: "Ada", want: "Thanks, Ada!"},
		{name: "whitespace-padded name", input: " Ada ", want: "Thanks, Ada!"},
		{name: "empty name", input: "", want: "Thank you!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Thanks(tt.input)
			if got != tt.want {
				t.Errorf("Thanks(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
