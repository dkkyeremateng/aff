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
		{name: "named", input: "Ada", want: "Thanks, Ada!"},
		{name: "whitespace-padded", input: " Ada ", want: "Thanks, Ada!"},
		{name: "empty", input: "", want: "Thank you!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Thanks(test.input)
			if got != test.want {
				t.Errorf("Thanks(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}
