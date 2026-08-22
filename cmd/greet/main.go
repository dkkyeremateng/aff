package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dkkyeremateng/aff/greet"
)

// version is the version reported by --version. It is a var with a
// development default so a release build can override it with
// -ldflags "-X main.version=<value>", which is the form main already carried.
var version = "dev"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var name string
	fs.StringVar(&name, "name", "", "name to greet")
	fs.StringVar(&name, "n", "", "name to greet (shorthand)")
	showVersion := fs.Bool("version", false, "print the version and exit")

	fs.Usage = func() {}

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			printUsage(stdout)
			return 0
		}
		if bad := unknownFlagName(fs, args); bad != "" {
			fmt.Fprintf(stderr, "greet: unknown flag: --%s\n", bad)
		} else {
			fmt.Fprintf(stderr, "greet: %v\n", err)
		}
		fmt.Fprintln(stderr, "Run 'greet --help' for usage.")
		return 2
	}

	if *showVersion {
		fmt.Fprintf(stdout, "greet version %s\n", version)
		return 0
	}

	fmt.Fprintln(stdout, greet.Hello(name))
	return 0
}

func printUsage(stdout io.Writer) {
	fmt.Fprintln(stdout, "Usage: greet [flags]")
	fmt.Fprintln(stdout)
	fmt.Fprintln(stdout, "Flags:")
	fmt.Fprintln(stdout, " -n, --name string name to greet")
	fmt.Fprintln(stdout, "  --version       print the version and exit")
	fmt.Fprintln(stdout, "  --help          show this help")
}

func unknownFlagName(fs *flag.FlagSet, args []string) string {
	for _, arg := range args {
		if len(arg) < 2 || arg[0] != '-' || arg == "--" {
			continue
		}
		name := strings.TrimLeft(arg, "-")
		if eq := strings.IndexByte(name, '='); eq >= 0 {
			name = name[:eq]
		}
		if name == "" || name == "help" || name == "h" {
			continue
		}
		if fs.Lookup(name) == nil {
			return name
		}
	}
	return ""
}
