package gitexec

import "context"

// Command executes an arbitrary Git subcommand with provided arguments.
func Command(dir, name string, arg ...string) ([]byte, error) {
	return CommandContext(context.Background(), dir, name, arg...)
}

// CommandContext executes an arbitrary Git subcommand with the provided context and arguments.
func CommandContext(ctx context.Context, dir, name string, arg ...string) ([]byte, error) {
	return run(execGit(ctx, dir, append([]string{name}, arg...)))
}
