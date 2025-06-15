package gitexec

import (
	"errors"
	"os/exec"
)

func run(cmd *exec.Cmd) ([]byte, error) {
	if cmd.Dir == "" {
		return nil, errors.New("missing command working directory")
	}

	withSysProcAttr(cmd)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, nil
}
