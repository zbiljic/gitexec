// Code generated for gitexec. DO NOT EDIT.
//
// git-reset - Reset current HEAD to the specified state
//
// Reference: https://git-scm.com/docs/git-reset
package gitexec

import (
	"fmt"
	"os/exec"
)

type ResetOptions struct {
	CmdDir string

	// -q
	// --quiet
	// Be quiet, only report errors.
	Quiet bool
	// -p
	// --patch
	// Interactively select hunks in the difference between the index and <tree-ish> to be applied in reverse to the index.
	Patch bool
	// --soft
	// Does not touch the index file or the working tree at all, but resets the head to <commit>. This leaves all your changed files as 'Changes to be committed'.
	Soft bool
	// --mixed
	// Resets the index but not the working tree (i.e., the changed files are preserved but not marked for commit) and reports what has not been updated. This is the default action.
	Mixed bool
	// --hard
	// Resets the index and working tree. Any changes to tracked files in the working tree since <commit> are discarded.
	Hard bool
	// --merge
	// Resets the index and updates the files in the working tree that are different between <commit> and HEAD, but keeps those which are different between the index and working tree.
	Merge bool
	// --keep
	// Resets index entries and updates files in the working tree that are different between <commit> and HEAD. If a file is different between <commit> and HEAD has local changes, reset is aborted.
	Keep bool
	// --no-recurse-submodules
	// Do not recursively reset the working tree of any active submodule according to the commit recorded in the superproject.
	NoRecurseSubmodules bool
	// --recurse-submodules
	// Recursively reset the working tree of all active submodules according to the commit recorded in the superproject.
	RecurseSubmodules bool
	// --refresh
	// Refresh the index after a mixed reset. Enabled by default.
	Refresh bool
	// --no-refresh
	// Do not refresh the index after a mixed reset.
	NoRefresh bool
	// --pathspec-from-file=<file>
	// Pathspec is passed in <file> instead of commandline args. If <file> is exactly - then standard input is used.
	PathspecFromFile string
	// --pathspec-file-nul
	// Only meaningful with --pathspec-from-file. Pathspec elements are separated with NUL character.
	PathspecFileNul bool
	// [<tree-ish>]
	// Resets the index entries for all <paths> to their state at <tree-ish>
	TreeIsh string
	// --
	// Do not interpret any more arguments as options.
	DoNotInterpretMoreArgumentsAsOptions bool
	// <pathspec>...
	// Limits the paths affected by the operation.
	Pathspec []string
	// [<commit>]
	// Resetting it to the tree of <commit>.
	Commit string
}

func ResetCmd(opts *ResetOptions) *exec.Cmd {
	args := []string{"reset"}

	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Patch {
		args = append(args, "--patch")
	}
	if opts.Soft {
		args = append(args, "--soft")
	}
	if opts.Mixed {
		args = append(args, "--mixed")
	}
	if opts.Hard {
		args = append(args, "--hard")
	}
	if opts.Merge {
		args = append(args, "--merge")
	}
	if opts.Keep {
		args = append(args, "--keep")
	}
	if opts.NoRecurseSubmodules {
		args = append(args, "--no-recurse-submodules")
	}
	if opts.RecurseSubmodules {
		args = append(args, "--recurse-submodules")
	}
	if opts.Refresh {
		args = append(args, "--refresh")
	}
	if opts.NoRefresh {
		args = append(args, "--no-refresh")
	}
	if opts.PathspecFromFile != "" {
		args = append(args, fmt.Sprintf("--pathspec-from-file=%s", opts.PathspecFromFile))
	}
	if opts.PathspecFileNul {
		args = append(args, "--pathspec-file-nul")
	}
	if opts.TreeIsh != "" {
		args = append(args, opts.TreeIsh)
	}
	if opts.DoNotInterpretMoreArgumentsAsOptions {
		args = append(args, "--")
	}
	if opts.Pathspec != nil {
		args = append(args, opts.Pathspec...)
	}
	if opts.Commit != "" {
		args = append(args, opts.Commit)
	}

	return execGit(opts.CmdDir, args)
}

func Reset(opts *ResetOptions) ([]byte, error) {
	return run(ResetCmd(opts))
}
