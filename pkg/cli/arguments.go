package cli

import (
	"flag"
	"fmt"
)

var (
	Template string
	Output   string
	Data     StringArray
	help     bool
)

func init() {
	Data = StringArray{}
	flag.Usage = func() {
		fmt.Println("Usage: gotmpl [options] -tmpl {URL,Path} -out {Path}")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.StringVar(&Template, "tmpl", "", "Template to use {URL,Path} or stdin.")
	flag.StringVar(&Output, "out", "", "Render output {Path} or stdout.")
	flag.Var(&Data, "data", "Data to use {URL,Path,JSON}. Multiple allowed.")
	flag.BoolVar(&help, "help", false, "Show this help message.")
}

func ParseArguments(args []string) bool {
	flag.CommandLine.Parse(args[1:])

	if help {
		flag.Usage()

		return true
	}

	return false
}
