package clients

import (
	"os/exec"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	var command = exec.Cmd{
		Path: CliName,
		Args: []string{"--help"},
	}

	// Expect no error
	_, err := ExecuteCommand(command)
	if err != nil {
		t.Error(err)
	}

	command.Args = append(command.Args, "-ulla")

	// Expect an error "unknown shorthand flag: 'u' in -ulla"
	_, err = ExecuteCommand(command)
	if err == nil {
		t.Error(err)
	}
}
