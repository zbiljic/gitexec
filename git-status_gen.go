// Code generated for gitexec. DO NOT EDIT.
//
// git-status - Show the working tree status
//
// Reference: https://git-scm.com/docs/git-status
package gitexec

import (
	"fmt"
	"os/exec"
)

type StatusOptions struct {
	CmdDir string

	// -s
	// --short
	// Give the output in the short-format.
	Short bool
	// -b
	// --branch
	// Show the branch and tracking info even in short-format.
	Branch bool
	// --show-stash
	// Show the number of entries currently stashed away.
	ShowStash bool
	// --porcelain[=<version>]
	// Give the output in an easy-to-parse format for scripts. This is similar to the short output, but will remain stable across Git versions and regardless of user configuration.
	Porcelain bool
	// --long
	// Give the output in the long-format. This is the default.
	Long bool
	// -v
	// --verbose
	// In addition to the names of files that have been changed, also show the textual changes that are staged to be committed.
	Verbose bool
	// -u[<mode>]
	// --untracked-files[=<mode>]
	// Show untracked files. The mode parameter is used to specify the handling of untracked files.
	UntrackedFiles bool
	// --ignore-submodules[=<when>]
	// Ignore changes to submodules when looking for changes.
	IgnoreSubmodules bool
	// --ignored[=<mode>]
	// Show ignored files as well. The mode parameter is used to specify the handling of ignored files.
	Ignored bool
	// -z
	// Terminate entries with NUL character instead of LF.
	Z bool
	// --column[=<options>]
	// Display untracked files in columns.
	Column bool
	// --no-column
	// Do not display untracked files in columns.
	NoColumn bool
	// --ahead-behind
	// Display or do not display detailed ahead/behind counts for the branch relative to its upstream branch.
	AheadBehind bool
	// --no-ahead-behind
	// Do not display detailed ahead/behind counts for the branch relative to its upstream branch.
	NoAheadBehind bool
	// --renames
	// Turn on rename detection regardless of user configuration.
	Renames bool
	// --no-renames
	// Turn off rename detection regardless of user configuration.
	NoRenames bool
	// --find-renames[=<n>]
	// Turn on rename detection, optionally setting the similarity threshold.
	FindRenames uint64
}

func StatusCmd(opts *StatusOptions) *exec.Cmd {
	args := []string{"status"}

	if opts.Short {
		args = append(args, "--short")
	}
	if opts.Branch {
		args = append(args, "--branch")
	}
	if opts.ShowStash {
		args = append(args, "--show-stash")
	}
	if opts.Porcelain {
		args = append(args, "--porcelain")
	}
	if opts.Long {
		args = append(args, "--long")
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.UntrackedFiles {
		args = append(args, "--untracked-files")
	}
	if opts.IgnoreSubmodules {
		args = append(args, "--ignore-submodules")
	}
	if opts.Ignored {
		args = append(args, "--ignored")
	}
	if opts.Z {
		args = append(args, "-z")
	}
	if opts.Column {
		args = append(args, "--column")
	}
	if opts.NoColumn {
		args = append(args, "--no-column")
	}
	if opts.AheadBehind {
		args = append(args, "--ahead-behind")
	}
	if opts.NoAheadBehind {
		args = append(args, "--no-ahead-behind")
	}
	if opts.Renames {
		args = append(args, "--renames")
	}
	if opts.NoRenames {
		args = append(args, "--no-renames")
	}
	if opts.FindRenames > 0 {
		args = append(args, fmt.Sprintf("--find-renames=%d", opts.FindRenames))
	}

	return execGit(opts.CmdDir, args)
}

func Status(opts *StatusOptions) ([]byte, error) {
	return run(StatusCmd(opts))
}
