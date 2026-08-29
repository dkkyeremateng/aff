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
	os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var name string
	fs.StringVar(&name, "name", "", "name to greet")
	fs.StringVar(&name, "n", "", "name to greet (shorthand)")
	var style string
	fs.StringVar(&style, "style", "hello", "greeting style")
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

	styleFuncs := map[string]func(string) string{
		"hello":    greet.Hello,
		"farewell": greet.Farewell,
		"thanks":   greet.Thanks,
		"welcome":  greet.Welcome,
		"congrats": greet.Congrats,
		"salute":   greet.Salute,
		"cheer":    greet.Cheer,
	}
	greetFn, ok := styleFuncs[style]
	if !ok {
		fmt.Fprintf(stderr, "greet: unknown style: %q (choose from hello, farewell, thanks, welcome, congrats, salute, cheer)\n", style)
		fmt.Fprintln(stderr, "Run 'greet --help' for usage.")
		return 2
	}

	nameProvided := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "name" || f.Name == "n" {
			nameProvided = true
		}
	})

	if !nameProvided && !stdinIsTerminal(stdin) {
		data, err := io.ReadAll(stdin)
		if err != nil {
			fmt.Fprintf(stderr, "greet: %v\n", err)
			return 1
		}
		name = strings.TrimSpace(string(data))
	}

	fmt.Fprintln(stdout, greetFn(name))
	return 0
}

func stdinIsTerminal(r io.Reader) bool {
	f, ok := r.(*os.File)
	if !ok {
		return false
	}
	fi, err := f.Stat()
	if err != nil {
		return true
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func printUsage(stdout io.Writer) {
	fmt.Fprintln(stdout, "Usage: greet [flags]")
	fmt.Fprintln(stdout)
	fmt.Fprintln(stdout, "Flags:")
	fmt.Fprintln(stdout, " -n, --name string name to greet")
	fmt.Fprintln(stdout, " --style string    greeting style (hello, farewell, thanks, welcome, congrats, salute, cheer; default hello)")
	fmt.Fprintln(stdout, " --version print the version and exit")
	fmt.Fprintln(stdout, " --help show this help")
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
