// Code generated for gitexec. DO NOT EDIT.
//
// git-clone - Clone a repository into a new directory
//
// Reference: https://git-scm.com/docs/git-clone
package gitexec

import (
	"fmt"
	"os/exec"
)

type CloneOptions struct {
	CmdDir string

	// --template=<template-directory>
	// Specify the directory from which templates will be used.
	Template string
	// -l
	// --local
	// When the repository to clone from is on a local machine, this flag bypasses the normal 'Git aware' transport mechanism and clones the repository by making a copy of HEAD and everything under objects and refs directories.
	Local bool
	// -s
	// --shared
	// Set up the new repository to share the object database with the cloned repository.
	Shared bool
	// --no-hardlinks
	// Force the cloning process from a repository on a local filesystem to copy the files under the .git/objects directory instead of using hardlinks.
	NoHardlinks bool
	// -q
	// --quiet
	// Operate quietly. Progress is not reported to the standard error stream.
	Quiet bool
	// -v
	// --verbose
	// Run verbosely. Does not affect the reporting of progress status to the standard error stream.
	Verbose bool
	// -n
	// --no-checkout
	// No checkout of HEAD is performed after the clone is complete.
	NoCheckout bool
	// --bare
	// Make a bare Git repository. That is, instead of creating <directory> and placing the administrative files in <directory>/.git, make the <directory> itself the $GIT_DIR.
	Bare bool
	// --mirror
	// Set up a mirror of the source repository. This implies --bare.
	Mirror bool
	// -o <name>
	// --origin <name>
	// Instead of using the remote name origin to keep track of the upstream repository, use <name>.
	Origin string
	// -b <name>
	// --branch <name>
	// Instead of pointing the newly created HEAD to the branch pointed to by the cloned repository's HEAD, point to <name> branch instead.
	Branch string
	// -u <upload-pack>
	// --upload-pack <upload-pack>
	// When given, and the repository to clone from is accessed via ssh, this specifies a non-default path for the command run on the other end.
	UploadPack string
	// --reference <repository>
	// If the reference repository is on the local machine, automatically setup .git/objects/info/alternates to obtain objects from the reference repository.
	Reference string
	// --reference-if-able <repository>
	// If the reference repository is on the local machine, automatically setup .git/objects/info/alternates to obtain objects from the reference repository, but only if the reference repository exists.
	ReferenceIfAble string
	// --dissociate
	// Borrow the objects from reference repositories specified with the --reference options only to reduce network transfer, and stop borrowing from them after a clone is made by making necessary local copies of borrowed objects.
	Dissociate bool
	// --separate-git-dir=<git-dir>
	// Instead of placing the cloned repository where it is supposed to be, place the cloned repository at the specified directory, then make a filesystem-agnostic Git symbolic link to there.
	SeparateGitDir string
	// --depth <depth>
	// Create a shallow clone with a history truncated to the specified number of commits.
	Depth string
	// --shallow-since=<date>
	// Create a shallow clone with a history after the specified time.
	ShallowSince string
	// --shallow-exclude=<revision>
	// Create a shallow clone with a history, excluding commits reachable from a specified remote branch or tag.
	ShallowExclude string
	// --single-branch
	// Clone only the history leading to the tip of a single branch, either specified by the --branch option or the primary branch remote's HEAD points at.
	SingleBranch bool
	// --no-single-branch
	// Override the --single-branch option to fetch the histories near the tips of all branches.
	NoSingleBranch bool
	// --no-tags
	// Don't clone any tags, and set remote.<remote>.tagOpt=--no-tags in the config file.
	NoTags bool
	// --recurse-submodules[=<pathspec>]
	// After the clone is created, initialize and clone submodules within based on the provided pathspec.
	RecurseSubmodules string
	// --shallow-submodules
	// All submodules which are cloned will be shallow with a depth of 1.
	ShallowSubmodules bool
	// --no-shallow-submodules
	// Disable the --shallow-submodules option.
	NoShallowSubmodules bool
	// --remote-submodules
	// All submodules which are cloned will use the status of the submodule's remote-tracking branch to update the submodule.
	RemoteSubmodules bool
	// --no-remote-submodules
	// Use the status of the submodule's head to update the submodule.
	NoRemoteSubmodules bool
	// -j <n>
	// --jobs <n>
	// The number of submodules fetched at the same time.
	Jobs uint64
	// --sparse
	// Initialize the sparse-checkout file so the working directory starts with only the files in the root of the repository.
	Sparse bool
	// --filter=<filter-spec>
	// Use the partial clone feature and request that the server sends a subset of reachable objects according to the specified filter.
	Filter string
	// --also-filter-submodules
	// When cloning with --filter, also apply the partial clone filter to any submodules in the repository.
	AlsoFilterSubmodules bool
	// --reject-shallow
	// Fail if the source repository is a shallow repository.
	RejectShallow bool
	// --no-reject-shallow
	// Overrides the --reject-shallow option.
	NoRejectShallow bool
	// --bundle-uri=<uri>
	// Before fetching from the remote, fetch a bundle from the given URI and unbundle the data into the local repository.
	BundleUri string
}

func CloneCmd(opts *CloneOptions) *exec.Cmd {
	args := []string{"clone"}

	if opts.Template != "" {
		args = append(args, fmt.Sprintf("--template=%s", opts.Template))
	}
	if opts.Local {
		args = append(args, "--local")
	}
	if opts.Shared {
		args = append(args, "--shared")
	}
	if opts.NoHardlinks {
		args = append(args, "--no-hardlinks")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.NoCheckout {
		args = append(args, "--no-checkout")
	}
	if opts.Bare {
		args = append(args, "--bare")
	}
	if opts.Mirror {
		args = append(args, "--mirror")
	}
	if opts.Origin != "" {
		args = append(args, "--origin")
		args = append(args, opts.Origin)
	}
	if opts.Branch != "" {
		args = append(args, "--branch")
		args = append(args, opts.Branch)
	}
	if opts.UploadPack != "" {
		args = append(args, "--upload-pack")
		args = append(args, opts.UploadPack)
	}
	if opts.Reference != "" {
		args = append(args, "--reference")
		args = append(args, opts.Reference)
	}
	if opts.ReferenceIfAble != "" {
		args = append(args, "--reference-if-able")
		args = append(args, opts.ReferenceIfAble)
	}
	if opts.Dissociate {
		args = append(args, "--dissociate")
	}
	if opts.SeparateGitDir != "" {
		args = append(args, fmt.Sprintf("--separate-git-dir=%s", opts.SeparateGitDir))
	}
	if opts.Depth != "" {
		args = append(args, "--depth")
		args = append(args, opts.Depth)
	}
	if opts.ShallowSince != "" {
		args = append(args, fmt.Sprintf("--shallow-since=%s", opts.ShallowSince))
	}
	if opts.ShallowExclude != "" {
		args = append(args, fmt.Sprintf("--shallow-exclude=%s", opts.ShallowExclude))
	}
	if opts.SingleBranch {
		args = append(args, "--single-branch")
	}
	if opts.NoSingleBranch {
		args = append(args, "--no-single-branch")
	}
	if opts.NoTags {
		args = append(args, "--no-tags")
	}
	if opts.RecurseSubmodules != "" {
		args = append(args, fmt.Sprintf("--recurse-submodules=%s", opts.RecurseSubmodules))
	}
	if opts.ShallowSubmodules {
		args = append(args, "--shallow-submodules")
	}
	if opts.NoShallowSubmodules {
		args = append(args, "--no-shallow-submodules")
	}
	if opts.RemoteSubmodules {
		args = append(args, "--remote-submodules")
	}
	if opts.NoRemoteSubmodules {
		args = append(args, "--no-remote-submodules")
	}
	if opts.Jobs > 0 {
		args = append(args, fmt.Sprintf("--jobs %d", opts.Jobs))
	}
	if opts.Sparse {
		args = append(args, "--sparse")
	}
	if opts.Filter != "" {
		args = append(args, fmt.Sprintf("--filter=%s", opts.Filter))
	}
	if opts.AlsoFilterSubmodules {
		args = append(args, "--also-filter-submodules")
	}
	if opts.RejectShallow {
		args = append(args, "--reject-shallow")
	}
	if opts.NoRejectShallow {
		args = append(args, "--no-reject-shallow")
	}
	if opts.BundleUri != "" {
		args = append(args, fmt.Sprintf("--bundle-uri=%s", opts.BundleUri))
	}

	return execGit(opts.CmdDir, args)
}

func Clone(opts *CloneOptions) ([]byte, error) {
	return run(CloneCmd(opts))
}
