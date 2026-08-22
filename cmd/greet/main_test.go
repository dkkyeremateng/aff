package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunGreetingGoesToStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--name", "Ada"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run() = %d, want 0", code)
	}
	if got, want := stdout.String(), "Hello, Ada!\n"; got != want {
		t.Errorf("stdout = %q, want %q", got, want)
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunDefaultGreetingGoesToStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run(nil, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run() = %d, want 0", code)
	}
	if got, want := stdout.String(), "Hello there!\n"; got != want {
		t.Errorf("stdout = %q, want %q", got, want)
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunVersionGoesToStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--version"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run() = %d, want 0", code)
	}
	if got, want := stdout.String(), "greet version "+version+"\n"; got != want {
		t.Errorf("stdout = %q, want %q", got, want)
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunHelpGoesToStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--help"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run() = %d, want 0", code)
	}
	if !strings.Contains(stdout.String(), "Usage: greet") {
		t.Errorf("stdout = %q, want usage text", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunUnknownFlagGoesToStderr(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--bogus"}, &stdout, &stderr)
	if code != 2 {
		t.Fatalf("run() = %d, want 2", code)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), "flag provided but not defined: -bogus") {
		t.Errorf("stderr = %q, want unknown flag diagnostic", stderr.String())
	}
	if !strings.Contains(stderr.String(), "Usage: greet") {
		t.Errorf("stderr = %q, want usage text", stderr.String())
	}
}

func TestRunUnexpectedArgumentGoesToStderr(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"extra"}, &stdout, &stderr)
	if code != 2 {
		t.Fatalf("run() = %d, want 2", code)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), `unexpected argument "extra"`) {
		t.Errorf("stderr = %q, want unexpected argument diagnostic", stderr.String())
	}
}
