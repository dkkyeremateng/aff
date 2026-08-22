package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		stdin           string
		wantCode        int
		wantStdout      string
		stdoutContains  []string
		stderrContains  []string
		wantStderrEmpty bool
		wantStderrOnly  bool
	}{
		{
			name:       "name flag",
			args:       []string{"--name", "Ada"},
			wantCode:   0,
			wantStdout: "Hello, Ada!\n",
		},
		{
			name:       "name shorthand",
			args:       []string{"-n", "Ada"},
			wantCode:   0,
			wantStdout: "Hello, Ada!\n",
		},
		{
			name:       "name shorthand with equals",
			args:       []string{"-n=Ada"},
			wantCode:   0,
			wantStdout: "Hello, Ada!\n",
		},
		{
			name:            "no args",
			args:            nil,
			wantCode:        0,
			wantStdout:      "Hello there!\n",
			wantStderrEmpty: true,
		},
		{
			name:            "no args empty slice",
			args:            []string{},
			wantCode:        0,
			wantStdout:      "Hello there!\n",
			wantStderrEmpty: true,
		},
		{
			name:            "stdin name",
			args:            nil,
			stdin:           "Ada\n",
			wantCode:        0,
			wantStdout:      "Hello, Ada!\n",
			wantStderrEmpty: true,
		},
		{
			name:            "stdin name padded",
			args:            nil,
			stdin:           " Ada \n",
			wantCode:        0,
			wantStdout:      "Hello, Ada!\n",
			wantStderrEmpty: true,
		},
		{
			name:            "stdin whitespace only",
			args:            nil,
			stdin:           " \n\t",
			wantCode:        0,
			wantStdout:      "Hello there!\n",
			wantStderrEmpty: true,
		},
		{
			name:            "flag wins over stdin",
			args:            []string{"--name", "Ada"},
			stdin:           "Bob\n",
			wantCode:        0,
			wantStdout:      "Hello, Ada!\n",
			wantStderrEmpty: true,
		},
		{
			name:            "empty name flag ignores stdin",
			args:            []string{"--name", ""},
			stdin:           "Bob\n",
			wantCode:        0,
			wantStdout:      "Hello there!\n",
			wantStderrEmpty: true,
		},
		{
			name:           "version ignores stdin",
			args:           []string{"--version"},
			stdin:          "Ada\n",
			wantCode:       0,
			stdoutContains: []string{"greet version dev"},
		},
		{
			name:           "version flag",
			args:           []string{"--version"},
			wantCode:       0,
			stdoutContains: []string{"greet version dev"},
		},
		{
			name:           "version wins over name",
			args:           []string{"--version", "--name", "Ada"},
			wantCode:       0,
			stdoutContains: []string{"greet version dev"},
		},
		{
			name:           "version wins over shorthand",
			args:           []string{"--version", "-n", "Ada"},
			wantCode:       0,
			stdoutContains: []string{"greet version dev"},
		},
		{
			name:           "unknown flag",
			args:           []string{"--frobnicate"},
			wantCode:       2,
			stderrContains: []string{"unknown flag: --frobnicate", "--help"},
			wantStderrOnly: true,
		},
		{
			name:           "unknown flag with value",
			args:           []string{"--frobnicate=1"},
			wantCode:       2,
			stderrContains: []string{"unknown flag: --frobnicate", "--help"},
			wantStderrOnly: true,
		},
		{
			name:     "help flag",
			args:     []string{"--help"},
			wantCode: 0,
			wantStdout: "Usage: greet [flags]\n" +
				"\n" +
				"Flags:\n" +
				" -n, --name string name to greet\n" +
				" --version         print the version and exit\n" +
				" --help            show this help\n",
			stdoutContains: []string{"-n", "--name", "--version"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, strings.NewReader(tt.stdin), &stdout, &stderr)
			if code != tt.wantCode {
				t.Errorf("run(%q) exit code = %d, want %d", tt.args, code, tt.wantCode)
			}
			if tt.wantStdout != "" && stdout.String() != tt.wantStdout {
				t.Errorf("run(%q) stdout = %q, want %q", tt.args, stdout.String(), tt.wantStdout)
			}
			for _, want := range tt.stdoutContains {
				if !strings.Contains(stdout.String(), want) {
					t.Errorf("run(%q) stdout = %q, want it to contain %q", tt.args, stdout.String(), want)
				}
			}
			for _, want := range tt.stderrContains {
				if !strings.Contains(stderr.String(), want) {
					t.Errorf("run(%q) stderr = %q, want it to contain %q", tt.args, stderr.String(), want)
				}
			}
			if tt.wantStderrEmpty && stderr.String() != "" {
				t.Errorf("run(%q) stderr = %q, want empty", tt.args, stderr.String())
			}
			if tt.wantStderrOnly {
				if stdout.String() != "" {
					t.Errorf("run(%q) stdout = %q, want empty", tt.args, stdout.String())
				}
				if strings.Contains(stderr.String(), "Usage of") {
					t.Errorf("run(%q) stderr = %q, contains default usage dump", tt.args, stderr.String())
				}
			}
		})
	}
}

func TestStdinIsTerminalNonFileReader(t *testing.T) {
	if stdinIsTerminal(strings.NewReader("Ada\n")) {
		t.Error("stdinIsTerminal(strings.Reader) = true, want false")
	}
}
