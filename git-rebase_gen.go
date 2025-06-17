// Code generated for gitexec. DO NOT EDIT.
//
// git-rebase - Reapply commits on top of another base tip
//
// Reference: https://git-scm.com/docs/git-rebase
package gitexec

import (
	"fmt"
	"os/exec"
)

type RebaseOptions struct {
	CmdDir string

	// --continue
	// Restart the rebasing process after having resolved a merge conflict.
	Continue bool
	// --skip
	// Restart the rebasing process by skipping the current patch.
	Skip bool
	// --abort
	// Abort the rebase operation and reset HEAD to the original branch.
	Abort bool
	// --quit
	// Abort the rebase operation but HEAD is not reset back to the original branch.
	Quit bool
	// --edit-todo
	// Edit the todo list during an interactive rebase.
	EditTodo bool
	// --show-current-patch
	// Show the current patch in an interactive rebase or when stopped.
	ShowCurrentPatch bool
	// --onto <newbase>
	// Starting point at which to create the new commits. If not specified, the starting point is <upstream>.
	Onto string
	// --keep-base
	// Set the starting point to the merge base of <upstream> and <branch>.
	KeepBase bool
	// --apply
	// Use applying strategies to rebase (calling git-am internally).
	Apply bool
	// --empty={drop
	// keep
	// ask}
	// How to handle commits that become empty (default: drop). 'keep' keeps empty commits, 'drop' drops them, 'ask' prompts for each.
	Empty string
	// --no-keep-empty
	// Do not keep commits that start empty before the rebase.
	NoKeepEmpty bool
	// --keep-empty
	// Keep commits that start empty before the rebase.
	KeepEmpty bool
	// --reapply-cherry-picks
	// Reapply all clean cherry-picks of any upstream commit instead of preemptively dropping them.
	ReapplyCherryPicks bool
	// --no-reapply-cherry-picks
	// Do not reapply all clean cherry-picks of any upstream commit instead of preemptively dropping them.
	NoReapplyCherryPicks bool
	// -m
	// --merge
	// Use merging strategies to rebase (default).
	Merge bool
	// -s <strategy>
	// --strategy=<strategy>
	// Use the given merge strategy. Implies --merge.
	Strategy string
	// -X <strategy-option>
	// --strategy-option=<strategy-option>
	// Pass the <strategy-option> through to the merge strategy.
	StrategyOption string
	// --rerere-autoupdate
	// Allow the rerere mechanism to update the index with the result of auto-conflict resolution.
	RerereAutoupdate bool
	// --no-rerere-autoupdate
	// Do not allow the rerere mechanism to update the index with the result of auto-conflict resolution.
	NoRerereAutoupdate bool
	// -S[<keyid>]
	// --gpg-sign[=<keyid>]
	// GPG-sign commits. The keyid argument is optional and defaults to the committer identity.
	GpgSign string
	// --no-gpg-sign
	// Do not GPG-sign the resulting merge commit. Useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.
	NoGpgSign bool
	// -q
	// --quiet
	// Be quiet. Implies --no-stat.
	Quiet bool
	// -v
	// --verbose
	// Be verbose. Implies --stat.
	Verbose bool
	// --stat
	// Show a diffstat of what changed upstream since the last rebase.
	Stat bool
	// -n
	// --no-stat
	// Do not show a diffstat as part of the rebase process.
	NoStat bool
	// --no-verify
	// Bypass the pre-rebase hook.
	NoVerify bool
	// --verify
	// Allows the pre-rebase hook to run, which is the default.
	Verify bool
	// -C<n>
	// Ensure at least <n> lines of surrounding context match before and after each change.
	Cn uint64
	// -f
	// --no-ff
	// --force-rebase
	// Force the rebase even if the current branch is a descendant of the commit you are rebasing onto.
	ForceRebase bool
	// --fork-point
	// Use reflog to find a better common ancestor between <upstream> and <branch>.
	ForkPoint bool
	// --no-fork-point
	// Do not use reflog to find a better common ancestor between <upstream> and <branch>.
	NoForkPoint bool
	// --ignore-whitespace
	// Ignore whitespace differences when trying to reconcile differences.
	IgnoreWhitespace bool
	// --whitespace=<option>
	// Ignore whitespace differences when trying to reconcile differences.
	Whitespace string
	// --committer-date-is-author-date
	// Use the author date of the commit being rebased as the committer date.
	CommitterDateIsAuthorDate bool
	// --ignore-date
	// --reset-author-date
	// Use the current time as the author date of the rebased commit.
	ResetAuthorDate bool
	// --signoff
	// Add a Signed-off-by trailer to all the rebased commits.
	Signoff bool
	// -i
	// --interactive
	// Make a list of commits to be rebased, let the user edit the list before proceeding. This mode can also be used to split commits.
	Interactive bool
	// --rebase-merges[=mode]
	// Preserve merge commits during rebase (mode can be 'rebase-cousins' or 'no-rebase-cousins').
	RebaseMerges string
	// --no-rebase-merges
	// Do not preserve merge commits during rebase.
	NoRebaseMerges bool
	// -x <cmd>
	// --exec <cmd>
	// Append 'exec <cmd>' after each line creating a commit in the final history.
	Exec string
	// --root
	// Rebase all commits reachable from <branch>, instead of limiting them with an <upstream>.
	Root bool
	// --autosquash
	// Automatically squash commits with specially formatted messages into previous commits.
	Autosquash bool
	// --no-autosquash
	// Do not automatically squash commits with specially formatted messages into previous commits.
	NoAutosquash bool
	// --autostash
	// Automatically create a temporary stash entry before the operation begins.
	Autostash bool
	// --no-autostash
	// Do not automatically create a temporary stash entry before the operation begins.
	NoAutostash bool
	// --reschedule-failed-exec
	// Automatically reschedule exec commands that failed.
	RescheduleFailedExec bool
	// --no-reschedule-failed-exec
	// Do not automatically reschedule exec commands that failed.
	NoRescheduleFailedExec bool
	// --update-refs
	// Automatically force-update any branches that point to commits that are being rebased.
	UpdateRefs bool
	// --no-update-refs
	// Do not automatically force-update any branches that point to commits that are being rebased.
	NoUpdateRefs bool
	// <upstream>
	// Upstream branch to compare against. Defaults to the configured upstream for the current branch.
	Upstream string
	// <branch>
	// Working branch; defaults to HEAD.
	Branch string
}

func RebaseCmd(opts *RebaseOptions) *exec.Cmd {
	args := []string{"rebase"}

	if opts.Continue {
		args = append(args, "--continue")
	}
	if opts.Skip {
		args = append(args, "--skip")
	}
	if opts.Abort {
		args = append(args, "--abort")
	}
	if opts.Quit {
		args = append(args, "--quit")
	}
	if opts.EditTodo {
		args = append(args, "--edit-todo")
	}
	if opts.ShowCurrentPatch {
		args = append(args, "--show-current-patch")
	}
	if opts.Onto != "" {
		args = append(args, "--onto")
		args = append(args, opts.Onto)
	}
	if opts.KeepBase {
		args = append(args, "--keep-base")
	}
	if opts.Apply {
		args = append(args, "--apply")
	}
	if opts.Empty != "" {
		args = append(args, fmt.Sprintf("--empty={drop,keep,ask}%s", opts.Empty))
	}
	if opts.NoKeepEmpty {
		args = append(args, "--no-keep-empty")
	}
	if opts.KeepEmpty {
		args = append(args, "--keep-empty")
	}
	if opts.ReapplyCherryPicks {
		args = append(args, "--reapply-cherry-picks")
	}
	if opts.NoReapplyCherryPicks {
		args = append(args, "--no-reapply-cherry-picks")
	}
	if opts.Merge {
		args = append(args, "--merge")
	}
	if opts.Strategy != "" {
		args = append(args, fmt.Sprintf("--strategy=%s", opts.Strategy))
	}
	if opts.StrategyOption != "" {
		args = append(args, fmt.Sprintf("--strategy-option=%s", opts.StrategyOption))
	}
	if opts.RerereAutoupdate {
		args = append(args, "--rerere-autoupdate")
	}
	if opts.NoRerereAutoupdate {
		args = append(args, "--no-rerere-autoupdate")
	}
	if opts.GpgSign != "" {
		args = append(args, fmt.Sprintf("--gpg-sign=%s", opts.GpgSign))
	}
	if opts.NoGpgSign {
		args = append(args, "--no-gpg-sign")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.Stat {
		args = append(args, "--stat")
	}
	if opts.NoStat {
		args = append(args, "--no-stat")
	}
	if opts.NoVerify {
		args = append(args, "--no-verify")
	}
	if opts.Verify {
		args = append(args, "--verify")
	}
	if opts.Cn > 0 {
		args = append(args, fmt.Sprintf("-C%d", opts.Cn))
	}
	if opts.ForceRebase {
		args = append(args, "--force-rebase")
	}
	if opts.ForkPoint {
		args = append(args, "--fork-point")
	}
	if opts.NoForkPoint {
		args = append(args, "--no-fork-point")
	}
	if opts.IgnoreWhitespace {
		args = append(args, "--ignore-whitespace")
	}
	if opts.Whitespace != "" {
		args = append(args, fmt.Sprintf("--whitespace=%s", opts.Whitespace))
	}
	if opts.CommitterDateIsAuthorDate {
		args = append(args, "--committer-date-is-author-date")
	}
	if opts.ResetAuthorDate {
		args = append(args, "--reset-author-date")
	}
	if opts.Signoff {
		args = append(args, "--signoff")
	}
	if opts.Interactive {
		args = append(args, "--interactive")
	}
	if opts.RebaseMerges != "" {
		args = append(args, fmt.Sprintf("--rebase-merges=%s", opts.RebaseMerges))
	}
	if opts.NoRebaseMerges {
		args = append(args, "--no-rebase-merges")
	}
	if opts.Exec != "" {
		args = append(args, "--exec")
		args = append(args, opts.Exec)
	}
	if opts.Root {
		args = append(args, "--root")
	}
	if opts.Autosquash {
		args = append(args, "--autosquash")
	}
	if opts.NoAutosquash {
		args = append(args, "--no-autosquash")
	}
	if opts.Autostash {
		args = append(args, "--autostash")
	}
	if opts.NoAutostash {
		args = append(args, "--no-autostash")
	}
	if opts.RescheduleFailedExec {
		args = append(args, "--reschedule-failed-exec")
	}
	if opts.NoRescheduleFailedExec {
		args = append(args, "--no-reschedule-failed-exec")
	}
	if opts.UpdateRefs {
		args = append(args, "--update-refs")
	}
	if opts.NoUpdateRefs {
		args = append(args, "--no-update-refs")
	}
	if opts.Upstream != "" {
		args = append(args, opts.Upstream)
	}
	if opts.Branch != "" {
		args = append(args, opts.Branch)
	}

	return execGit(opts.CmdDir, args)
}

func Rebase(opts *RebaseOptions) ([]byte, error) {
	return run(RebaseCmd(opts))
}
