// Code generated for gitexec. DO NOT EDIT.
//
// git-symbolic-ref - Read, create, or delete symbolic refs
//
// Reference: https://git-scm.com/docs/git-symbolic-ref
package gitexec

import "os/exec"

type SymbolicRefOptions struct {
	CmdDir string

	// -d
	// --delete
	// Delete the symbolic ref <name>.
	Delete bool
	// -q
	// --quiet
	// Do not issue an error message if the <name> is not a symbolic ref but a detached HEAD; instead exit with status 1 silently.
	Quiet bool
	// --short
	// When showing a symbolic ref, try to shorten it as a relative ref.
	Short bool
	// --recurse
	// When showing the value of <name> as a symbolic ref, if <name> refers to another symbolic ref, follow such a chain of symbolic refs until the result no longer points at a symbolic ref.
	Recurse bool
	// --no-recurse
	// When showing the value of <name> as a symbolic ref, if <name> refers to another symbolic ref, stop after dereferencing only a single level of symbolic ref.
	NoRecurse bool
	// -m <reason>
	// Update the reflog for <name> with <reason>. This is valid only when creating or updating a symbolic ref.
	M string
	// <name>
	// The name of the symbolic ref to read, create, or delete. For example HEAD.
	Name string
	// <ref>
	// The target of the symbolic ref. For example refs/heads/master.
	Ref string
}

func SymbolicRefCmd(opts *SymbolicRefOptions) *exec.Cmd {
	args := []string{"symbolic-ref"}

	if opts.Delete {
		args = append(args, "--delete")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Short {
		args = append(args, "--short")
	}
	if opts.Recurse {
		args = append(args, "--recurse")
	}
	if opts.NoRecurse {
		args = append(args, "--no-recurse")
	}
	if opts.M != "" {
		args = append(args, "-m")
		args = append(args, opts.M)
	}
	if opts.Name != "" {
		args = append(args, opts.Name)
	}
	if opts.Ref != "" {
		args = append(args, opts.Ref)
	}

	return execGit(opts.CmdDir, args)
}

func SymbolicRef(opts *SymbolicRefOptions) ([]byte, error) {
	return run(SymbolicRefCmd(opts))
}
