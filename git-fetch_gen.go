// Code generated for gitexec. DO NOT EDIT.
//
// git-fetch - Download objects and refs from another repository
//
// Reference: https://git-scm.com/docs/git-fetch
package gitexec

import (
	"fmt"
	"os/exec"
)

type FetchOptions struct {
	CmdDir string

	// --all
	// Fetch all remotes, except for the ones that has the remote.<name>.skipFetchAll configuration variable set.
	All bool
	// --no-all
	// Override the configuration variable fetch.all, which when set to true fetches all remotes.
	NoAll bool
	// -a
	// --append
	// Append ref names and object names of fetched refs to the existing contents of .git/FETCH_HEAD.
	Append bool
	// --atomic
	// Use an atomic transaction to update local refs. Either all refs are updated, or on error, no refs are updated.
	Atomic bool
	// --depth=<depth>
	// Limit fetching to the specified number of commits from the tip of each remote branch history.
	Depth string
	// --deepen=<depth>
	// Similar to --depth, except it specifies the number of commits from the current shallow boundary instead of from the tip of each remote branch history.
	Deepen string
	// --shallow-since=<date>
	// Deepen or shorten the history of a shallow repository to include all reachable commits after <date>.
	ShallowSince string
	// --shallow-exclude=<ref>
	// Deepen or shorten the history of a shallow repository to exclude commits reachable from a specified remote branch or tag.
	ShallowExclude string
	// --unshallow
	// If the source repository is complete, convert a shallow repository to a complete one by removing all the limitations imposed by shallow repositories.
	Unshallow bool
	// --update-shallow
	// Update .git/shallow for the current shallow repository to include what is fetched.
	UpdateShallow bool
	// --negotiation-tip=<commit|glob>
	// By default, Git will report all reachable commits from its local refs when asking a remote what commits it needs to complete a fetch.
	NegotiationTip string
	// --negotiate-only
	// Do not actually fetch, but only print the hashes of commits the client has and the server needs.
	NegotiateOnly bool
	// --dry-run
	// Show what would be done, without making any changes.
	DryRun bool
	// --write-fetch-head
	// Write the list of remote refs fetched in the FETCH_HEAD file directly under $GIT_DIR.
	WriteFetchHead bool
	// --no-write-fetch-head
	// Tells Git not to write the updated refs in the FETCH_HEAD file.
	NoWriteFetchHead bool
	// -f
	// --force
	// When git fetch is used with <src>:<dst> refspec it may refuse to update the local branch which is not an ancestor of the remote branch. This option overrides that check.
	Force bool
	// -k
	// --keep
	// Keep downloaded pack.
	Keep bool
	// --multiple
	// Allow several <repository> and <group> arguments to be specified. No <refspec>s may be specified.
	Multiple bool
	// --auto-maintenance
	// Run 'git maintenance run --auto' after fetching.
	AutoMaintenance bool
	// --no-auto-maintenance
	// Don't run 'git maintenance run --auto' after fetching.
	NoAutoMaintenance bool
	// --auto-gc
	// Run 'git gc --auto' after fetching.
	AutoGc bool
	// --no-auto-gc
	// Don't run 'git gc --auto' after fetching.
	NoAutoGc bool
	// --write-commit-graph
	// Write a commit-graph after fetching.
	WriteCommitGraph bool
	// --no-write-commit-graph
	// Don't write a commit-graph after fetching.
	NoWriteCommitGraph bool
	// --prefetch
	// Modify the configured refspec to place all refs into the refs/prefetch/ namespace.
	Prefetch bool
	// -p
	// --prune
	// Before fetching, remove any remote-tracking references that no longer exist on the remote.
	Prune bool
	// -P
	// --prune-tags
	// Before fetching, remove any local tags that no longer exist on the remote if --prune is enabled.
	PruneTags bool
	// -n
	// --no-tags
	// By default, tags are fetched automatically if the remote helps us to determine which objects would be transfered if we fetch from the remote branches.
	NoTags bool
	// --refetch
	// Instead of negotiating with the server to avoid transferring commits and associated objects that are already present locally, this option fetches all objects as a fresh clone would.
	Refetch bool
	// --refmap=<refspec>
	// When fetching refs listed on the command line, use the specified refspec to map the refs to remote-tracking branches.
	Refmap string
	// -t
	// --tags
	// Fetch all tags from the remote (i.e., fetch remote tags refs/tags/* into local tags with the same name).
	Tags bool
	// --recurse-submodules[=(yes|on-demand|no)]
	// This option controls if and under what conditions new commits of populated submodules should be fetched too.
	RecurseSubmodules string
	// -j
	// --jobs=<n>
	// Number of parallel children to be used for fetching submodules.
	Jobs uint64
	// --no-recurse-submodules
	// Disable recursive fetching of submodules.
	NoRecurseSubmodules bool
	// -u
	// --set-upstream
	// If the remote is fetched successfully, add upstream (tracking) reference.
	SetUpstream bool
	// --submodule-prefix=<path>
	// Prepend <path> to paths printed in informative messages such as 'Fetching submodule foo'.
	SubmodulePrefix string
	// --recurse-submodules-default=[yes|on-demand]
	// This option sets the default value for the --recurse-submodules option.
	RecurseSubmodulesDefault string
	// -u
	// --update-head-ok
	// By default git fetch refuses to update the head which corresponds to the current branch.
	UpdateHeadOk bool
	// --upload-pack <upload-pack>
	// When given, and the repository to fetch from is accessed via ssh, this specifies a non-default path for the command run on the other end.
	UploadPack string
	// -q
	// --quiet
	// Progress status is reported on the standard error stream by default, suppress it.
	Quiet bool
	// -v
	// --verbose
	// Be verbose.
	Verbose bool
	// --progress
	// Progress status is reported on the standard error stream by default, but this flag forces progress status even if the error stream is not directed to a terminal.
	Progress bool
	// -o
	// --server-option=<option>
	// Transmit the given string to the server when communicating using protocol version 2.
	ServerOption string
	// --show-forced-updates
	// By default, git fetch prints forced updates with a different preceding symbol than normal updates.
	ShowForcedUpdates bool
	// --no-show-forced-updates
	// Do not print forced updates.
	NoShowForcedUpdates bool
	// -4
	// --ipv4
	// Use IPv4 addresses only, ignoring IPv6 addresses.
	Ipv4 bool
	// -6
	// --ipv6
	// Use IPv6 addresses only, ignoring IPv4 addresses.
	Ipv6 bool
	// <repository>
	// The "remote" repository that is the source of a fetch or pull operation.
	Repository string
	// <group>
	// A name referring to a list of repositories as the value of remotes.<group> in the configuration file.
	Group string
	// <refspec>
	// Specifies which refs to fetch and which local refs to update.
	Refspec string
	// --stdin
	// Read refspecs from stdin in addition to those provided as arguments.
	Stdin bool
}

func FetchCmd(opts *FetchOptions) *exec.Cmd {
	args := []string{"fetch"}

	if opts.All {
		args = append(args, "--all")
	}
	if opts.NoAll {
		args = append(args, "--no-all")
	}
	if opts.Append {
		args = append(args, "--append")
	}
	if opts.Atomic {
		args = append(args, "--atomic")
	}
	if opts.Depth != "" {
		args = append(args, fmt.Sprintf("--depth=%s", opts.Depth))
	}
	if opts.Deepen != "" {
		args = append(args, fmt.Sprintf("--deepen=%s", opts.Deepen))
	}
	if opts.ShallowSince != "" {
		args = append(args, fmt.Sprintf("--shallow-since=%s", opts.ShallowSince))
	}
	if opts.ShallowExclude != "" {
		args = append(args, fmt.Sprintf("--shallow-exclude=%s", opts.ShallowExclude))
	}
	if opts.Unshallow {
		args = append(args, "--unshallow")
	}
	if opts.UpdateShallow {
		args = append(args, "--update-shallow")
	}
	if opts.NegotiationTip != "" {
		args = append(args, fmt.Sprintf("--negotiation-tip=%s", opts.NegotiationTip))
	}
	if opts.NegotiateOnly {
		args = append(args, "--negotiate-only")
	}
	if opts.DryRun {
		args = append(args, "--dry-run")
	}
	if opts.WriteFetchHead {
		args = append(args, "--write-fetch-head")
	}
	if opts.NoWriteFetchHead {
		args = append(args, "--no-write-fetch-head")
	}
	if opts.Force {
		args = append(args, "--force")
	}
	if opts.Keep {
		args = append(args, "--keep")
	}
	if opts.Multiple {
		args = append(args, "--multiple")
	}
	if opts.AutoMaintenance {
		args = append(args, "--auto-maintenance")
	}
	if opts.NoAutoMaintenance {
		args = append(args, "--no-auto-maintenance")
	}
	if opts.AutoGc {
		args = append(args, "--auto-gc")
	}
	if opts.NoAutoGc {
		args = append(args, "--no-auto-gc")
	}
	if opts.WriteCommitGraph {
		args = append(args, "--write-commit-graph")
	}
	if opts.NoWriteCommitGraph {
		args = append(args, "--no-write-commit-graph")
	}
	if opts.Prefetch {
		args = append(args, "--prefetch")
	}
	if opts.Prune {
		args = append(args, "--prune")
	}
	if opts.PruneTags {
		args = append(args, "--prune-tags")
	}
	if opts.NoTags {
		args = append(args, "--no-tags")
	}
	if opts.Refetch {
		args = append(args, "--refetch")
	}
	if opts.Refmap != "" {
		args = append(args, fmt.Sprintf("--refmap=%s", opts.Refmap))
	}
	if opts.Tags {
		args = append(args, "--tags")
	}
	if opts.RecurseSubmodules != "" {
		args = append(args, fmt.Sprintf("--recurse-submodules=%s", opts.RecurseSubmodules))
	}
	if opts.Jobs > 0 {
		args = append(args, fmt.Sprintf("--jobs=%d", opts.Jobs))
	}
	if opts.NoRecurseSubmodules {
		args = append(args, "--no-recurse-submodules")
	}
	if opts.SetUpstream {
		args = append(args, "--set-upstream")
	}
	if opts.SubmodulePrefix != "" {
		args = append(args, fmt.Sprintf("--submodule-prefix=%s", opts.SubmodulePrefix))
	}
	if opts.RecurseSubmodulesDefault != "" {
		args = append(args, fmt.Sprintf("--recurse-submodules-default=%s", opts.RecurseSubmodulesDefault))
	}
	if opts.UpdateHeadOk {
		args = append(args, "--update-head-ok")
	}
	if opts.UploadPack != "" {
		args = append(args, "--upload-pack")
		args = append(args, opts.UploadPack)
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.Progress {
		args = append(args, "--progress")
	}
	if opts.ServerOption != "" {
		args = append(args, fmt.Sprintf("--server-option=%s", opts.ServerOption))
	}
	if opts.ShowForcedUpdates {
		args = append(args, "--show-forced-updates")
	}
	if opts.NoShowForcedUpdates {
		args = append(args, "--no-show-forced-updates")
	}
	if opts.Ipv4 {
		args = append(args, "--ipv4")
	}
	if opts.Ipv6 {
		args = append(args, "--ipv6")
	}
	if opts.Stdin {
		args = append(args, "--stdin")
	}
	if opts.Repository != "" {
		args = append(args, opts.Repository)
	}
	if opts.Group != "" {
		args = append(args, opts.Group)
	}
	if opts.Refspec != "" {
		args = append(args, opts.Refspec)
	}

	return execGit(opts.CmdDir, args)
}

func Fetch(opts *FetchOptions) ([]byte, error) {
	return run(FetchCmd(opts))
}
