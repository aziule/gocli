# gocli

gocli helps you implement CLI commands for your app by providing:
- A common interface for your CLI commands
- A handler to handle 'em all


How is that useful?
- One single handler for every CLI command
- No need to recreate the wheel every time you want your app to serve CLI commands
- Easy debugging
- Easy to use
- Fast implementation of new CLI commands

## Getting Started

[![GoDoc](https://godoc.org/github.com/Aziule/gocli?status.svg)](https://godoc.org/github.com/Aziule/gocli)
[![Go Report Card](https://goreportcard.com/badge/github.com/Aziule/gocli)](https://goreportcard.com/report/github.com/Aziule/gocli)
[![Build Status](https://travis-ci.org/Aziule/gocli.svg?branch=master)](https://travis-ci.org/aziule/gocli)
[![License](http://img.shields.io/:license-mit-blue.svg)](LICENSE)

### Prerequisites

Developed with Go 1.8+ (installation instructions [here](https://golang.org/doc/install))

### Installing

`go get -u github.com/aziule/gocli`

### Usage

First, create a new CLI command and make it implement the `Command` interface:

```golang
// RunCommand contains the variables set by SetFlags (if any)
// that will be used within Execute()
type RunCommand struct {
    configFilePath string
}

// Usage describes how our command works
func (c *RunCommand) Usage() string {
    return `run [-config=./config.json]:
    The description of what the command is doing`
}

// Execute runs the command
func (c *RunCommand) Execute(f *flag.FlagSet) error {
    // Do whatever you need here
    fmt.Println(c.configFilePath)
    
    return nil
}

// FlagSet defines what flags to set for the command
func (c *RunCommand) SetFlags(f *flag.FlagSet) {
    // This will store the flag's value to the command's structure
    f.StringVar(&c.configFilePath, "config", "config.json", "Config file path")
}

// Name returns the command's name to be invoked from the CLI
func (c *RunCommand) Name() string {
    return "run"
}
```

Then, create a new `CliHandler` and register the previous CLI command:

```golang
func main() {
    cliHandler := new gocli.CliHandler()
    cliHandler.RegisterCommand(&RunCommand{})
    err := cliHandler.Handle()
    
    if err != nil {
        // Handle the error your way
    }
}
```

**Running the command**
```bash
# From the dev environment
go run ./*.go <command> -config /path/to/config.json

# Using the binary
/path/to/myapp <command> -config /path/to/config.json
```

For example:

```bash
/path/to/myapp run -config /path/to/config.json
# "Here is the config: config.json"
```

**Viewing the usage**
```bash
/path/to/myapp
# COMMANDS:
# run [-config=./config.json]:
#     The description of what the command is doing
```

## Contributing

Please feel free to report any issue or improvement suggestion.

## Authors

* **William Claude** - *Developer* - [Github](https://github.com/aziule)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
