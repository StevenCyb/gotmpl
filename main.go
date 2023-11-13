package main

import (
	"encoding/json"
	"fmt"
	"gotmpl/pkg/cli"
	"gotmpl/pkg/engine"
	"gotmpl/pkg/input"
	"gotmpl/pkg/output"
	"io"
	"os"
	"strings"
)

var stdin io.Reader = os.Stdin

func main() {
	if cli.ParseArguments(os.Args) {
		return
	}

	// Process input
	var template []byte
	if err := processInput("template", cli.Template, &template); err != nil {
		fmt.Println(err.Error())
		return //os.Exit(1)
	}

	var data []interface{}
	for i, dataSource := range cli.Data {
		var rawData []byte
		if err := processInput("data", dataSource, &rawData); err != nil {
			fmt.Println(err.Error())
			return //os.Exit(1)
		}

		var dataParsed interface{}
		if err := json.Unmarshal(rawData, &dataParsed); err != nil {
			fmt.Printf("Failed to parse data source [%d]: %s", i, err.Error())
			return //os.Exit(1)
		}

		data = append(data, dataParsed)
	}

	// Process template
	executor := engine.New()
	executor.SetTemplate(template)
	executor.SetData(data)
	result, err := executor.Execute()
	if err != nil {
		fmt.Printf("Failed to execute template: %s", err.Error())

		return
	}

	// Perform output
	if cli.Output != "" {
		if err := output.ToFile(cli.Output, result); err != nil {
			fmt.Printf("Failed to write output to file: %s", err.Error())
			return //os.Exit(1)
		}

		return
	}

	if err := output.ToWriter(os.Stdout, result); err != nil {
		fmt.Printf("Failed to write output to stdout: %s", err.Error())
		return //os.Exit(1)
	}
}

func processInput(name, from string, to *[]byte) error {
	var err error

	switch {
	case strings.HasPrefix(from, "http://") || strings.HasPrefix(from, "https://"):
		// example `go run main.go -tmpl https://www.google.com`
		if *to, err = input.FromWeb(from); err != nil {
			return fmt.Errorf("Reading %s from web failed: %s\n", name, err.Error())
		}
	case strings.HasPrefix(from, "{") || strings.HasPrefix(from, "["):
		// example `go run main.go -tmpl '{"test": "test"}'`
		if _, err = json.Marshal(from); err != nil {
			return fmt.Errorf("Reading %s from json input failed: %s\n", name, err.Error())
		}
		*to = []byte(from)
	case from != "":
		// example `go run main.go -tmpl main.go`
		if *to, err = input.FromFile(from); err != nil {
			return fmt.Errorf("Reading %s from file failed: %s\n", name, err.Error())
		}
	case from == "":
		// example `echo "test" | go run main.go`
		if *to, err = input.FromReader(stdin); err != nil {
			return fmt.Errorf("Reading %s from stdin failed: %s\n", name, err.Error())
		}
	default:
		return fmt.Errorf("unknown %s source %s\n", name, from)
	}

	return nil
}
