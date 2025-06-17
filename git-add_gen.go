// Code generated for gitexec. DO NOT EDIT.
//
// git-add - Add file contents to the index
//
// Reference: https://git-scm.com/docs/git-add
package gitexec

import (
	"fmt"
	"os/exec"
)

type AddOptions struct {
	CmdDir string

	// -n
	// --dry-run
	// Don't actually add the file(s), just show if they exist and/or will be ignored.
	DryRun bool
	// -v
	// --verbose
	// Be verbose.
	Verbose bool
	// -f
	// --force
	// Allow adding otherwise ignored files.
	Force bool
	// --sparse
	// Allow updating index entries outside of the sparse-checkout cone.
	Sparse bool
	// -i
	// --interactive
	// Add modified contents in the working tree interactively to the index. Optional path arguments may be supplied to limit operation to a subset of the working tree.
	Interactive bool
	// -p
	// --patch
	// Interactively choose hunks of patch between the index and the work tree and add them to the index.
	Patch bool
	// -e
	// --edit
	// Open the diff vs. the index in an editor and let the user edit it. After the editor was closed, adjust the hunk headers and apply the patch to the index.
	Edit bool
	// -u
	// --update
	// Update the index just where it already has an entry matching <pathspec>. This removes as well as modifies index entries to match the working tree, but adds no new files.
	Update bool
	// -A
	// --all
	// --no-ignore-removal
	// Update the index not only where the working tree has a file matching <pathspec> but also where the index already has an entry. This adds, modifies, and removes index entries to match the working tree.
	All bool
	// --no-all
	// --ignore-removal
	// Update the index by adding new files that are unknown to the index and files modified in the working tree, but ignore files that have been removed from the working tree.
	NoAll bool
	// -N
	// --intent-to-add
	// Record only the fact that the path will be added later. An entry for the path is placed in the index with no content.
	IntentToAdd bool
	// --refresh
	// Don't add the file(s), but only refresh their stat() information in the index.
	Refresh bool
	// --ignore-errors
	// If some files could not be added because of errors indexing them, do not abort the operation, but continue adding the others.
	IgnoreErrors bool
	// --ignore-missing
	// This option can only be used together with --dry-run. By using this option the user can check if any of the given files would be ignored, no matter if they are already present in the work tree or not.
	IgnoreMissing bool
	// --no-warn-embedded-repo
	// Suppress the warning when adding an embedded repository to the index without using git submodule add.
	NoWarnEmbeddedRepo bool
	// --renormalize
	// Apply the "clean" process freshly to all tracked files to forcibly add them again to the index. This option implies -u.
	Renormalize bool
	// --chmod=(+|-)x
	// Override the executable bit of the added files. The executable bit is only changed in the index, the files on disk are left unchanged.
	Chmod string
	// --pathspec-from-file=<file>
	// Pathspec is passed in <file> instead of commandline args. If <file> is exactly - then standard input is used.
	PathspecFromFile string
	// --pathspec-file-nul
	// Only meaningful with --pathspec-from-file. Pathspec elements are separated with NUL character.
	PathspecFileNul bool
	// --
	// Separate command-line options from the list of files.
	DoNotInterpretMoreArgumentsAsOptions bool
	// <pathspec>...
	// Files to add content from. Fileglobs (e.g. *.c) can be given to add all matching files.
	Pathspec []string
}

func AddCmd(opts *AddOptions) *exec.Cmd {
	args := []string{"add"}

	if opts.DryRun {
		args = append(args, "--dry-run")
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.Force {
		args = append(args, "--force")
	}
	if opts.Sparse {
		args = append(args, "--sparse")
	}
	if opts.Interactive {
		args = append(args, "--interactive")
	}
	if opts.Patch {
		args = append(args, "--patch")
	}
	if opts.Edit {
		args = append(args, "--edit")
	}
	if opts.Update {
		args = append(args, "--update")
	}
	if opts.All {
		args = append(args, "--all")
	}
	if opts.NoAll {
		args = append(args, "--no-all")
	}
	if opts.IntentToAdd {
		args = append(args, "--intent-to-add")
	}
	if opts.Refresh {
		args = append(args, "--refresh")
	}
	if opts.IgnoreErrors {
		args = append(args, "--ignore-errors")
	}
	if opts.IgnoreMissing {
		args = append(args, "--ignore-missing")
	}
	if opts.NoWarnEmbeddedRepo {
		args = append(args, "--no-warn-embedded-repo")
	}
	if opts.Renormalize {
		args = append(args, "--renormalize")
	}
	if opts.Chmod != "" {
		args = append(args, fmt.Sprintf("--chmod=%s", opts.Chmod))
	}
	if opts.PathspecFromFile != "" {
		args = append(args, fmt.Sprintf("--pathspec-from-file=%s", opts.PathspecFromFile))
	}
	if opts.PathspecFileNul {
		args = append(args, "--pathspec-file-nul")
	}
	if opts.DoNotInterpretMoreArgumentsAsOptions {
		args = append(args, "--")
	}
	if opts.Pathspec != nil {
		args = append(args, opts.Pathspec...)
	}

	return execGit(opts.CmdDir, args)
}

func Add(opts *AddOptions) ([]byte, error) {
	return run(AddCmd(opts))
}
