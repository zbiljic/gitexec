//go:build unix

package gitexec

import (
	"os/exec"
	"syscall"
)

func withSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
}

func withCancel(cmd *exec.Cmd) {
	if cmd.Cancel == nil {
		return
	}

	cmd.Cancel = func() error {
		if cmd.Process == nil {
			return nil
		}

		if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil && err != syscall.ESRCH {
			return err
		}

		return nil
	}
}
