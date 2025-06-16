// Code generated for gitexec. DO NOT EDIT.
//
// git-pull - Fetch from and integrate with another repository or a local branch
//
// Reference: https://git-scm.com/docs/git-pull
package gitexec

import (
	"fmt"
	"os/exec"
)

type PullOptions struct {
	CmdDir string

	// -q
	// --quiet
	// This is passed to both underlying git-fetch to squelch reporting of during transfer, and underlying git-merge to squelch output during merging.
	Quiet bool
	// -v
	// --verbose
	// Pass --verbose to git-fetch and git-merge.
	Verbose bool
	// --recurse-submodules[=(yes|on-demand|no)]
	// This option controls if new commits of populated submodules should be fetched, and if the working trees of active submodules should be updated.
	RecurseSubmodules string
	// --no-recurse-submodules
	// Disable recursive fetching of submodules.
	NoRecurseSubmodules bool
	// --commit
	// Perform the merge and commit the result. This option can be used to override --no-commit. Only useful when merging.
	Commit bool
	// --no-commit
	// Perform the merge and stop just before creating a merge commit, to give the user a chance to inspect and further tweak the merge result before committing. Note that fast-forward updates do not create a merge commit and therefore there is no way to stop those merges with --no-commit. Thus, if you want to ensure your branch is not changed or updated by the merge command, use --no-ff with --no-commit.
	NoCommit bool
	// -e
	// --edit
	// Invoke an editor before committing successful mechanical merge to further edit the auto-generated merge message, so that the user can explain and justify the merge.
	Edit bool
	// --no-edit
	// Accept the auto-generated merge message without invoking an editor. This is generally discouraged.
	NoEdit bool
	// --cleanup=<mode>
	// This option determines how the merge message will be cleaned up before committing. See git-commit[1] for more details. In addition, if the <mode> is given a value of scissors, scissors will be appended to MERGE_MSG before being passed on to the commit machinery in the case of a merge conflict.
	Cleanup string
	// --ff-only
	// Only update to the new history if there is no divergent local history. This is the default when no method for reconciling divergent histories is provided (via the --rebase=* flags).
	FfOnly bool
	// --ff
	// When merging rather than rebasing, specifies how a merge is handled when the merged-in history is already a descendant of the current history. With --ff, when possible resolve the merge as a fast-forward (only update the branch pointer to match the merged branch; do not create a merge commit). When not possible (when the merged-in history is not a descendant of the current history), create a merge commit.
	Ff bool
	// --no-ff
	// Create a merge commit in all cases, even when the merge could instead be resolved as a fast-forward.
	NoFf bool
	// -S[<keyid>]
	// --gpg-sign[=<keyid>]
	// GPG-sign the resulting merge commit. The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space.
	GpgSign string
	// --no-gpg-sign
	// Do not GPG-sign the resulting merge commit. Useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.
	NoGpgSign bool
	// --log[=<n>]
	// In addition to branch names, populate the log message with at most <n> actual commits that are being merged. See also git-fmt-merge-msg[1]. Only useful when merging.
	Log uint64
	// --no-log
	// Do not list one-line descriptions from the actual commits being merged. Only useful when merging.
	NoLog bool
	// --signoff
	// Add a Signed-off-by trailer by the committer at the end of the commit log message. The meaning of a signoff depends on the project. Consult the documentation or leadership of the project to understand how signoffs are used.
	Signoff bool
	// --no-signoff
	// Countermand an earlier --signoff option on the command line.
	NoSignoff bool
	// --stat
	// Show a diffstat at the end of the merge. The diffstat is also controlled by the configuration option merge.stat.
	Stat bool
	// -n
	// --no-stat
	// Do not show a diffstat at the end of the merge.
	NoStat bool
	// --squash
	// Produce the working tree and index state as if a real merge happened (except for the merge information), but do not actually make a commit, move the HEAD, or record $GIT_DIR/MERGE_HEAD. This allows you to create a single commit on top of the current branch whose effect is the same as merging another branch. With --squash, --commit is not allowed, and will fail. Only useful when merging.
	Squash bool
	// --no-squash
	// Perform the merge and commit the result. This option can be used to override --squash.
	NoSquash bool
	// --verify
	// Run the pre-merge and commit-msg hooks. This is the default. Only useful when merging.
	Verify bool
	// --no-verify
	// Bypass the pre-merge and commit-msg hooks. See also githooks[5]. Only useful when merging.
	NoVerify bool
	// -s <strategy>
	// --strategy=<strategy>
	// Use the given merge strategy; can be supplied more than once to specify them in the order they should be tried. If there is no -s option, a built-in list of strategies is used instead.
	Strategy string
	// -X <option>
	// --strategy-option=<option>
	// Pass merge strategy specific option through to the merge strategy.
	StrategyOption string
	// --verify-signatures
	// Verify that the GPG signature of the fetched objects is valid.
	VerifySignatures bool
	// --no-verify-signatures
	// Do not verify the GPG signature of the fetched objects.
	NoVerifySignatures bool
	// --autostash
	// Automatically create a temporary stash entry before the operation begins, and apply it after the operation ends. This means that you can run rebase on a dirty worktree. However, use with care: the final stash application after a successful rebase might result in non-trivial conflicts.
	Autostash bool
	// --no-autostash
	// Disable the --autostash option.
	NoAutostash bool
	// --allow-unrelated-histories
	// This option can be used to override this safety when merging histories of two projects that started their lives independently. Only useful when merging.
	AllowUnrelatedHistories bool
	// -r
	// --rebase[=false|true|merges|interactive]
	// When true, rebase the current branch on top of the upstream branch after fetching. If there is a remote-tracking branch corresponding to the upstream branch and the upstream branch was rebased since last fetched, the rebase uses that information to avoid rebasing non-local changes. When set to merges, rebase using the --rebase-merges option. When false, merge the current branch into the upstream branch.
	Rebase string
	// --no-rebase
	// This is shorthand for --rebase=false.
	NoRebase bool
	// --all
	// Fetch all remotes.
	All bool
	// --no-all
	// Override the configuration variable fetch.all, which when set to true fetches all remotes.
	NoAll bool
	// -a
	// --append
	// Append ref names and object names of fetched refs to the existing contents of .git/FETCH_HEAD. Without this option old data in .git/FETCH_HEAD will be overwritten.
	Append bool
	// --atomic
	// Use an atomic transaction to update local refs. Either all refs are updated, or on error, no refs are updated.
	Atomic bool
	// --depth=<depth>
	// Limit fetching to the specified number of commits from the tip of each remote branch history. If fetching to a shallow repository created by git clone with --depth=<depth> option (see git-clone[1]), deepen or shorten the history to the specified number of commits.
	Depth string
	// --deepen=<depth>
	// Similar to --depth, except it specifies the number of commits from the current shallow boundary instead of from the tip of each remote branch history.
	Deepen string
	// --shallow-since=<date>
	// Deepen or shorten the history of a shallow repository to include all reachable commits after <date>.
	ShallowSince string
	// --shallow-exclude=<revision>
	// Deepen or shorten the history of a shallow repository to exclude commits reachable from a specified remote branch or tag. This option can be specified multiple times.
	ShallowExclude string
	// --unshallow
	// If the source repository is complete, convert a shallow repository to a complete one, removing all the limitations imposed by shallow repositories.
	Unshallow bool
	// --update-shallow
	// By default when fetching from a shallow repository, git fetch refuses refs that require updating .git/shallow. This option updates .git/shallow and accept such refs.
	UpdateShallow bool
	// --negotiation-tip=<commit|glob>
	// By default, Git will report, to the server, refs reachable from all local branches and tags. This option allows you to specify a particular commit or glob pattern for commits that should be considered as tips for negotiation.
	NegotiationTip string
	// --negotiate-only
	// Do not fetch anything from the server, and instead print the ancestors of the provided --negotiation-tip=* arguments, which we have in common with the server.
	NegotiateOnly bool
	// --dry-run
	// Show what would be done, without making any changes.
	DryRun bool
	// --porcelain
	// Print the output to standard output in an easy-to-parse format for scripts.
	Porcelain bool
	// -f
	// --force
	// When git fetch is used with <rbranch>:<lbranch> refspec, it refuses to update the local branch <lbranch> unless the remote branch <rbranch> it fetches is a descendant of <lbranch>. This option overrides that check.
	Force bool
	// -k
	// --keep
	// Keep downloaded pack.
	Keep bool
	// --prefetch
	// Modify the configured refspec to place all refs into the refs/prefetch/ namespace.
	Prefetch bool
	// -p
	// --prune
	// Before fetching, remove any remote-tracking references that no longer exist on the remote.
	Prune bool
	// --no-tags
	// By default, tags that point at objects that are downloaded from the remote repository are fetched and stored locally. This option disables this automatic tag following. The default behavior for a remote may be specified with the remote.<name>.tagOpt setting. See git-config[1].
	NoTags bool
	// --refmap=<refspec>
	// When fetching refs listed on the command line, use the specified refspec (can be given more than once) to map the refs to remote-tracking branches, instead of the values of remote.*.fetch configuration variables for the remote repository. Providing an empty <refspec> to the --refmap option causes Git to ignore the configured refmaps and rely entirely on the refspecs supplied as command-line arguments. See section on "Configured Remote-tracking Branches" for details.
	Refmap string
	// -t
	// --tags
	// Fetch all tags from the remote (i.e., fetch remote tags refs/tags/* into local tags with the same name), in addition to whatever else would otherwise be fetched. Using this option alone does not subject tags to pruning, even if --prune is used (though tags may be pruned anyway if they are also the destination of a refspec explicitly mapped to them).
	Tags bool
	// -j <n>
	// --jobs=<n>
	// Number of parallel children to be used for fetching submodules. Defaults to the submodule.fetchJobs option. If the option is not set, it defaults to 1.
	Jobs uint64
	// --set-upstream
	// If the remote is fetched successfully, add upstream (tracking) reference, used by argument-less git-pull[1] and other commands.
	SetUpstream bool
	// -u <upload-pack>
	// --upload-pack <upload-pack>
	// When given, and the repository to clone from is accessed via ssh, this specifies a non-default path for the command run on the other end.
	UploadPack string
	// --progress
	// Progress status is reported on the standard error stream by default when it is attached to a terminal, unless -q is specified. This flag forces progress status even if the standard error stream is not directed to a terminal.
	Progress bool
	// -o <option>
	// --server-option=<option>
	// Transmit the given string to the server when communicating using protocol version 2. The given string must not contain a NUL or LF character. When multiple are given, they are all sent in order.
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
	// <refspec>
	// Specifies which refs to fetch and which local refs to update.
	Refspec string
}

func PullCmd(opts *PullOptions) *exec.Cmd {
	args := []string{"pull"}

	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.RecurseSubmodules != "" {
		args = append(args, fmt.Sprintf("--recurse-submodules=%s", opts.RecurseSubmodules))
	}
	if opts.NoRecurseSubmodules {
		args = append(args, "--no-recurse-submodules")
	}
	if opts.Commit {
		args = append(args, "--commit")
	}
	if opts.NoCommit {
		args = append(args, "--no-commit")
	}
	if opts.Edit {
		args = append(args, "--edit")
	}
	if opts.NoEdit {
		args = append(args, "--no-edit")
	}
	if opts.Cleanup != "" {
		args = append(args, fmt.Sprintf("--cleanup=%s", opts.Cleanup))
	}
	if opts.FfOnly {
		args = append(args, "--ff-only")
	}
	if opts.Ff {
		args = append(args, "--ff")
	}
	if opts.NoFf {
		args = append(args, "--no-ff")
	}
	if opts.GpgSign != "" {
		args = append(args, fmt.Sprintf("--gpg-sign=%s", opts.GpgSign))
	}
	if opts.NoGpgSign {
		args = append(args, "--no-gpg-sign")
	}
	if opts.Log > 0 {
		args = append(args, fmt.Sprintf("--log=%d", opts.Log))
	}
	if opts.NoLog {
		args = append(args, "--no-log")
	}
	if opts.Signoff {
		args = append(args, "--signoff")
	}
	if opts.NoSignoff {
		args = append(args, "--no-signoff")
	}
	if opts.Stat {
		args = append(args, "--stat")
	}
	if opts.NoStat {
		args = append(args, "--no-stat")
	}
	if opts.Squash {
		args = append(args, "--squash")
	}
	if opts.NoSquash {
		args = append(args, "--no-squash")
	}
	if opts.Verify {
		args = append(args, "--verify")
	}
	if opts.NoVerify {
		args = append(args, "--no-verify")
	}
	if opts.Strategy != "" {
		args = append(args, fmt.Sprintf("--strategy=%s", opts.Strategy))
	}
	if opts.StrategyOption != "" {
		args = append(args, fmt.Sprintf("--strategy-option=%s", opts.StrategyOption))
	}
	if opts.VerifySignatures {
		args = append(args, "--verify-signatures")
	}
	if opts.NoVerifySignatures {
		args = append(args, "--no-verify-signatures")
	}
	if opts.Autostash {
		args = append(args, "--autostash")
	}
	if opts.NoAutostash {
		args = append(args, "--no-autostash")
	}
	if opts.AllowUnrelatedHistories {
		args = append(args, "--allow-unrelated-histories")
	}
	if opts.Rebase != "" {
		args = append(args, fmt.Sprintf("--rebase=%s", opts.Rebase))
	}
	if opts.NoRebase {
		args = append(args, "--no-rebase")
	}
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
	if opts.Porcelain {
		args = append(args, "--porcelain")
	}
	if opts.Force {
		args = append(args, "--force")
	}
	if opts.Keep {
		args = append(args, "--keep")
	}
	if opts.Prefetch {
		args = append(args, "--prefetch")
	}
	if opts.Prune {
		args = append(args, "--prune")
	}
	if opts.NoTags {
		args = append(args, "--no-tags")
	}
	if opts.Refmap != "" {
		args = append(args, fmt.Sprintf("--refmap=%s", opts.Refmap))
	}
	if opts.Tags {
		args = append(args, "--tags")
	}
	if opts.Jobs > 0 {
		args = append(args, fmt.Sprintf("--jobs=%d", opts.Jobs))
	}
	if opts.SetUpstream {
		args = append(args, "--set-upstream")
	}
	if opts.UploadPack != "" {
		args = append(args, "--upload-pack")
		args = append(args, opts.UploadPack)
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
	if opts.Repository != "" {
		args = append(args, opts.Repository)
	}
	if opts.Refspec != "" {
		args = append(args, opts.Refspec)
	}

	return execGit(opts.CmdDir, args)
}

func Pull(opts *PullOptions) ([]byte, error) {
	return run(PullCmd(opts))
}
