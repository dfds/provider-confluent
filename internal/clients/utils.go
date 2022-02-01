package clients

import (
	"os"
	"os/exec"
)

// ExecuteCommand Execute command helper method
func ExecuteCommand(cmd exec.Cmd) ([]byte, error) {
	execCmd := exec.Command(cmd.Path, cmd.Args...) //nolint:gosec
	execCmd.Env = os.Environ()

	out, err := execCmd.CombinedOutput()

	if err != nil {
		return out, err
	}

	return out, err
}
