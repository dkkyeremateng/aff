package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunHelp(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"--help"}, &stdout, &stderr)
	if code != 0 {
		t.Errorf("run(--help) = %d, want %d", code, 0)
	}
	for _, want := range []string{"Usage: greet [flags]", "--help", "--name", "--version"} {
		if !strings.Contains(stdout.String(), want) {
			t.Errorf("stdout = %q, want it to contain %q", stdout.String(), want)
		}
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunVersion(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"--version"}, &stdout, &stderr)
	if code != 0 {
		t.Errorf("run(--version) = %d, want %d", code, 0)
	}
	if got, want := stdout.String(), "greet dev\n"; got != want {
		t.Errorf("stdout = %q, want %q", got, want)
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunGreets(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "no args", args: nil, want: "Hello there!\n"},
		{name: "separate name value", args: []string{"--name", "Ada"}, want: "Hello, Ada!\n"},
		{name: "inline name value", args: []string{"--name=Ada"}, want: "Hello, Ada!\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer

			code := run(tt.args, &stdout, &stderr)
			if code != 0 {
				t.Errorf("run(%q) = %d, want %d", tt.args, code, 0)
			}
			if got := stdout.String(); got != tt.want {
				t.Errorf("stdout = %q, want %q", got, tt.want)
			}
			if stderr.Len() != 0 {
				t.Errorf("stderr = %q, want empty", stderr.String())
			}
		})
	}
}

func TestRunUnknownFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"--bogus"}, &stdout, &stderr)
	if code != 2 {
		t.Errorf("run(--bogus) = %d, want %d", code, 2)
	}
	for _, want := range []string{"unknown flag: --bogus", "Usage: greet [flags]"} {
		if !strings.Contains(stderr.String(), want) {
			t.Errorf("stderr = %q, want it to contain %q", stderr.String(), want)
		}
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
}

func TestRunNameMissingValue(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"--name"}, &stdout, &stderr)
	if code != 2 {
		t.Errorf("run(--name) = %d, want %d", code, 2)
	}
	if want := "flag --name requires a value"; !strings.Contains(stderr.String(), want) {
		t.Errorf("stderr = %q, want it to contain %q", stderr.String(), want)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
}

func TestRunHelpBeatsOtherFlags(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"--version", "--help"}, &stdout, &stderr)
	if code != 0 {
		t.Errorf("run(--version --help) = %d, want %d", code, 0)
	}
	if want := "Usage: greet [flags]"; !strings.Contains(stdout.String(), want) {
		t.Errorf("stdout = %q, want it to contain %q", stdout.String(), want)
	}
}
