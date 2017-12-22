# gocli
(Yet Another) Go CLI management tool

[![GoDoc](https://godoc.org/github.com/Aziule/gocli?status.svg)](https://godoc.org/github.com/Aziule/gocli)

gocli provides everything you need in order to add CLI to an app.

## Usage

The first step is to create the available commands. To do so, create a `struct` for each command you want to be able to run, and make them
implement the `Command` interface.

For example:

```golang
// RunCommand is the command responsible for running our app
type RunCommand struct {
	configFilePath string
}

// Usage returns the usage text for the command
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

// FlagSet returns the command's flag set
func (c *RunCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.configFilePath, "config", "config.json", "Config file path")
}

// Name returns the command's name, to be used when invoking it from the cli
func (c *RunCommand) Name() string {
	return "run"
}
```

Then, create a `Handler` to register your commands and make it handle the commands:

```golang
func main() {
	cliHandler := gocli.NewHandler()
	cliHandler.RegisterCommand(&RunCommand{})
	err := cliHandler.Handle()

	if err != nil {
		panic(fmt.Sprintf("An error occurred when handling the command: %s", err))
	}
}
```

You are now able to run `go run ./*.go run`
