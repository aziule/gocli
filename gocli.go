// Package gocli provides everything you need in order to add CLI to an app.
// @todo: add a Writer for writing to the output
package gocli

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Command is the main interface for creating new commands to be used from the CLI
type Command interface {
	Name() string
	Execute(*flag.FlagSet) error
	SetFlags(*flag.FlagSet)
	Usage() string
}

// CliHandler is the handler responsible for storing and executing commands
type CliHandler struct {
	topLevelFlags *flag.FlagSet
	commands      []Command
}

// NewHandler creates a new CliHandler
func NewHandler() *CliHandler {
	handler := &CliHandler{
		topLevelFlags: flag.CommandLine,
	}

	handler.topLevelFlags.Usage = func() { handler.explain() }

	return handler
}

// RegisterCommand adds a command to the list of available commands within a given handler
func (h *CliHandler) RegisterCommand(command Command) {
	h.commands = append(h.commands, command)
}

// Handle handles a command identified by its name
func (h *CliHandler) Handle() error {
	if !flag.Parsed() {
		flag.Parse()
	}

	if h.topLevelFlags.NArg() < 1 {
		h.topLevelFlags.Usage()
		return errors.New("Invalid arguments passed")
	}

	name := h.topLevelFlags.Arg(0)

	for _, command := range h.commands {
		if command.Name() != name {
			continue
		}

		f := flag.NewFlagSet(name, flag.ContinueOnError)
		f.Usage = func() { fmt.Println(command.Usage()) }
		command.SetFlags(f)

		if err := f.Parse(h.topLevelFlags.Args()[1:]); err != nil {
			return fmt.Errorf("Could not parse flags: %s", err)
		}

		return command.Execute(f)
	}

	h.topLevelFlags.Usage()

	return fmt.Errorf("Command %s not found", name)
}

// explain explains to the user how to use the commands and what commands are available
func (h *CliHandler) explain() {
	fmt.Fprintln(os.Stderr, "COMMANDS:")

	for _, c := range h.commands {
		fmt.Fprintln(os.Stderr, c.Usage())
	}
}
