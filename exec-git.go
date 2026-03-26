package gitexec

import (
	"context"
	"os/exec"
)

func execGit(ctx context.Context, dir string, args []string) *exec.Cmd {
	var cmd *exec.Cmd
	if ctx == nil {
		cmd = exec.Command("git", args...)
	} else {
		cmd = exec.CommandContext(ctx, "git", args...)
	}
	cmd.Dir = dir
	return cmd
}
