// Code generated for gitexec. DO NOT EDIT.
//
// git-ls-remote - List references in a remote repository
//
// Reference: https://git-scm.com/docs/git-ls-remote
package gitexec

import (
	"fmt"
	"os/exec"
)

type LsRemoteOptions struct {
	CmdDir string

	// -b
	// --branches
	// Limit to only references stored in refs/heads. Note that --heads and -h are deprecated synonyms for --branches and -b.
	Branches bool
	// -t
	// --tags
	// Limit to only references stored in refs/tags. When both --branches and --tags are given, references from both refs/heads and refs/tags are displayed.
	Tags bool
	// --refs
	// Do not show peeled tags or pseudorefs like HEAD in the output.
	Refs bool
	// -q
	// --quiet
	// Do not print remote URL to stderr.
	Quiet bool
	// --upload-pack=<exec>
	// Specify the full path of git-upload-pack on the remote host. This allows listing references from repositories accessed via SSH and where the SSH daemon does not use the PATH configured by the user.
	UploadPack string
	// --exit-code
	// Exit with status "2" when no matching refs are found in the remote repository. Usually the command exits with status "0".
	ExitCode bool
	// --get-url
	// Expand the URL of the given remote repository taking into account any "url.<base>.insteadOf" config setting and exit without talking to the remote.
	GetUrl bool
	// --symref
	// In addition to the object pointed by it, show the underlying ref pointed by it when showing a symbolic ref. Currently, upload-pack only shows the symref HEAD.
	Symref bool
	// --sort=<key>
	// Sort based on the key given. Prefix - to sort in descending order. Supports "version:refname" or "v:refname". See git-for-each-ref[1] for more sort options.
	Sort string
	// -o <option>
	// --server-option=<option>
	// Transmit the given string to the server when communicating using protocol version 2. The given string must not contain a NUL or LF character. When multiple are given, they are all sent in order.
	ServerOption string
	// <repository>
	// The "remote" repository to query.
	Repository string
	// <patterns>...
	// When specified, only references matching one or more of the given patterns are displayed.
	Patterns []string
}

func LsRemoteCmd(opts *LsRemoteOptions) *exec.Cmd {
	args := []string{"ls-remote"}

	if opts.Branches {
		args = append(args, "--branches")
	}
	if opts.Tags {
		args = append(args, "--tags")
	}
	if opts.Refs {
		args = append(args, "--refs")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.UploadPack != "" {
		args = append(args, fmt.Sprintf("--upload-pack=%s", opts.UploadPack))
	}
	if opts.ExitCode {
		args = append(args, "--exit-code")
	}
	if opts.GetUrl {
		args = append(args, "--get-url")
	}
	if opts.Symref {
		args = append(args, "--symref")
	}
	if opts.Sort != "" {
		args = append(args, fmt.Sprintf("--sort=%s", opts.Sort))
	}
	if opts.ServerOption != "" {
		args = append(args, fmt.Sprintf("--server-option=%s", opts.ServerOption))
	}
	if opts.Repository != "" {
		args = append(args, opts.Repository)
	}
	if opts.Patterns != nil {
		args = append(args, opts.Patterns...)
	}

	return execGit(opts.CmdDir, args)
}

func LsRemote(opts *LsRemoteOptions) ([]byte, error) {
	return run(LsRemoteCmd(opts))
}
