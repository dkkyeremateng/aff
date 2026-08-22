// Command greet prints a greeting.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dkkyeremateng/aff/greet"
)

// version is the version reported by the --version flag. It defaults to a
// development value and may be overridden at build time with
// -ldflags "-X main.version=<value>".
var version = "dev"

func run(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(stderr)

	showVersion := fs.Bool("version", false, "print the version and exit")
	name := fs.String("name", "", "name to greet")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *showVersion {
		fmt.Fprintln(stdout, version)
		return nil
	}

	fmt.Fprintln(stdout, greet.Hello(*name))
	return nil
}

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		os.Exit(2)
	}
}
