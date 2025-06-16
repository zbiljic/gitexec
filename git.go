package gitexec

// Command executes an arbitrary Git subcommand with provided arguments.
func Command(dir, name string, arg ...string) ([]byte, error) {
	return run(execGit(dir, append([]string{name}, arg...)))
}
