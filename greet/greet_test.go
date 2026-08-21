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

func TestCongrats(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "direct name", input: "Ada", want: "Congrats, Ada!"},
		{name: "whitespace-padded name", input: " Ada ", want: "Congrats, Ada!"},
		{name: "empty name", input: "", want: "Congratulations!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Congrats(tt.input)
			if got != tt.want {
				t.Errorf("Congrats(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestGoodNight(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "direct name", input: "Ada", want: "Good night, Ada!"},
		{name: "whitespace-padded name", input: " Ada ", want: "Good night, Ada!"},
		{name: "empty name", input: "", want: "Good night!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GoodNight(tt.input)
			if got != tt.want {
				t.Errorf("GoodNight(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestWelcome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "direct name", input: "Ada", want: "Welcome, Ada!"},
		{name: "whitespace-padded name", input: " Ada ", want: "Welcome, Ada!"},
		{name: "empty name", input: "", want: "Welcome!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Welcome(tt.input)
			if got != tt.want {
				t.Errorf("Welcome(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
