package main

import (
	"bytes"
	"testing"
)

func TestVersionFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer

	if err := run([]string{"--version"}, &stdout, &stderr); err != nil {
		t.Fatalf("run(--version) returned error: %v", err)
	}

	want := version + "\n"
	if got := stdout.String(); got != want {
		t.Errorf("stdout = %q, want %q", got, want)
	}
	if got := stderr.String(); got != "" {
		t.Errorf("stderr = %q, want empty", got)
	}
}

func TestDefaultInvocation(t *testing.T) {
	var stdout, stderr bytes.Buffer

	if err := run(nil, &stdout, &stderr); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}

	want := "Hello there!\n"
	if got := stdout.String(); got != want {
		t.Errorf("stdout = %q, want %q", got, want)
	}
	if got := stderr.String(); got != "" {
		t.Errorf("stderr = %q, want empty", got)
	}
}

func TestNameFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer

	if err := run([]string{"--name", "Ada"}, &stdout, &stderr); err != nil {
		t.Fatalf("run(--name Ada) returned error: %v", err)
	}

	want := "Hello, Ada!\n"
	if got := stdout.String(); got != want {
		t.Errorf("stdout = %q, want %q", got, want)
	}
}
