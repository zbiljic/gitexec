package gitexec

import "os/exec"

func execGit(dir string, args []string) *exec.Cmd {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	return cmd
}
