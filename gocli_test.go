package gocli

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @todo: add missing tests and fix fmt.Println stuff

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

func TestRegisterCommand(t *testing.T) {
	handler := NewHandler()
	cmd := &RunCommand{}
	handler.RegisterCommand(cmd)

	assert.Equal(t, 1, len(handler.commands))
	assert.ObjectsAreEqual(handler.commands[0], cmd)
}

func TestHandleWithNoFlags(t *testing.T) {
	handler := NewHandler()
	cmd := &RunCommand{}
	handler.RegisterCommand(cmd)

	err := handler.Handle()
	assert.Equal(t, "Invalid arguments passed", err.Error())
}

func TestHandleUnexistingCommand(t *testing.T) {
	handler := NewHandler()
	cmd := &RunCommand{}
	handler.RegisterCommand(cmd)

	handler.topLevelFlags.Parse([]string{"impossibru"})

	err := handler.Handle()
	assert.Equal(t, "Command impossibru not found", err.Error())
}

func TestHandleCommandWithDefaultFlagsValues(t *testing.T) {
}

func TestHandleCommand(t *testing.T) {
}
