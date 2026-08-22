// Command greet prints a friendly greeting.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dkkyeremateng/aff/greet"
)

const version = "dev"

// usage writes the command usage text to w.
func usage(w io.Writer) {
	fmt.Fprintln(w, "Usage: greet [flags]")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Flags:")
	fmt.Fprintln(w, "  --help     print this help and exit")
	fmt.Fprintln(w, "  --name     name to greet")
	fmt.Fprintln(w, "  --version  print the version and exit")
}

// run parses args and executes the command, writing the greeting to stdout
// and diagnostics to stderr. It returns the process exit code.
func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {}
	showVersion := fs.Bool("version", false, "print the version and exit")
	name := fs.String("name", "", "name to greet")

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			usage(stdout)
			return 0
		}
		usage(stderr)
		return 2
	}

	if *showVersion {
		fmt.Fprintln(stdout, "greet version "+version)
		return 0
	}

	if fs.NArg() > 0 {
		fmt.Fprintf(stderr, "greet: unexpected argument %q\n", fs.Arg(0))
		usage(stderr)
		return 2
	}

	fmt.Fprintln(stdout, greet.Hello(*name))
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
