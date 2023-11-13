# gotmpl

`gotmpl` is a command-line interface (CLI) tool for template rendering. Whether you're working with templates from stdin, files, or web resources, `gotmpl` simplifies the process by allowing you to provide multiple data inputs through arguments, files, or URLs. The rendered output can be directed to stdout or saved into a specified file.

## Example
```bash
$ echo "Hello, {{.Name}}" | go run main.go -data '{"Name": "John"}'
Hello, John%     
```

## Usage
Currently supported arguments can be listed as follow.
```bash
$ ./gotmpl --help
Usage: gotmpl [options] -tmpl {URL,Path} -out {Path}
Options:
  -data value
        Data to use {URL,Path,JSON}. Multiple allowed.
  -help
        Show this help message.
  -out string
        Render output {Path} or stdout.
  -tmpl string
        Template to use {URL,Path} or stdin.
```

This cli tool uses the go text template engine. A documentation for that is available [here](https://pkg.go.dev/text/template).
In general, the tool can be used as follow:
1. provide a template
   * from a file like `-tmpl /some/path/mail_template.txt`
   * from a web resource like `-tmpl https://some.url/mail_template` (requires `http` or `https` prefix)
   * from std like `echo "template" | ...` (without `-tmpl` argument)
2. provide data *(for multiple data source see note below)*
   * from a JSON formatted file like `-tmpl /some/path/user_a.json`
   * from a JSON formatted web resource like `-tmpl https://some.url/user_a` (requires `http` or `https` prefix)
   * directly as JSON formatted string `-tmpl '{...}'` or `-tmpl '[...]'`
3. define the output
   * into a file like `-out result.txt`
   * stdout is used by default if `-out` not defined

**NOTE:** When using multiple data sources, they are placed in an array (keeping the order). 
For example the templates will look as follow multiple data sources:
```bash
echo "Name: {{(index . 0).Name}}, Age: {{(index . 1).Age}}" | ./gotmpl -data '{"Name": "Nina"}' -data '{"Age": 33}'
# OR
echo "Name: {{index . 0 \"Name\"}}, Age: {{index . 1 \"Age\"}}" | ./gotmpl -data '{"Name": "Nina"}' -data '{"Age": 33}'
```
With single data source:
```bash
echo "Name: {{.Name}}, Age: {{.Age}}" | ./gotmpl -data '{"Name": "Nina", "Age": 33}'
```

