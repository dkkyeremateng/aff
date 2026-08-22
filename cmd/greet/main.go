// Command greet prints a friendly greeting.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dkkyeremateng/aff/greet"
)

const usage = `Usage: greet [flags]

Flags:
  --help          show this help and exit
  --name <name>   name to greet (default: generic greeting)
  --version       print version and exit
`

var version = "dev"

func run(args []string, stdout, stderr io.Writer) int {
	name := ""
	showVersion := false

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--help":
			fmt.Fprint(stdout, usage)
			return 0
		case arg == "--version":
			showVersion = true
		case arg == "--name":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "greet: flag --name requires a value")
				fmt.Fprint(stderr, usage)
				return 2
			}
			i++
			name = args[i]
		case strings.HasPrefix(arg, "--name="):
			name = strings.TrimPrefix(arg, "--name=")
		case strings.HasPrefix(arg, "-"):
			fmt.Fprintf(stderr, "greet: unknown flag: %s\n", arg)
			fmt.Fprint(stderr, usage)
			return 2
		default:
			fmt.Fprintf(stderr, "greet: unexpected argument: %s\n", arg)
			fmt.Fprint(stderr, usage)
			return 2
		}
	}

	if showVersion {
		fmt.Fprintf(stdout, "greet %s\n", version)
		return 0
	}

	fmt.Fprintln(stdout, greet.Hello(name))
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
