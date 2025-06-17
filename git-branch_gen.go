// Code generated for gitexec. DO NOT EDIT.
//
// git-branch - List, create, or delete branches
//
// Reference: https://git-scm.com/docs/git-branch
package gitexec

import (
	"fmt"
	"os/exec"
)

type BranchOptions struct {
	CmdDir string

	// -d
	// --delete
	// Delete a branch. The branch must be fully merged in its upstream branch, or in HEAD if no upstream was set.
	Delete bool
	// -D
	// Shortcut for --delete --force.
	D bool
	// --create-reflog
	// Create the branch's reflog. This activates recording of all changes made to the branch ref, enabling use of date based sha1 expressions such as "<branchname>@{yesterday}".
	CreateReflog bool
	// -f
	// --force
	// Reset <branchname> to <startpoint>, even if <branchname> exists already. Without -f, git branch refuses to change an existing branch.
	Force bool
	// -m
	// --move
	// Move/rename a branch and the corresponding reflog.
	Move bool
	// -M
	// Shortcut for --move --force.
	M bool
	// -c
	// --copy
	// Copy a branch and the corresponding reflog.
	Copy bool
	// -C
	// Shortcut for --copy --force.
	C bool
	// --color[=<when>]
	// Color branches to highlight current, local, and remote branches.
	Color string
	// --no-color
	// Turn off branch coloring, even when the configuration file gives the default to color output.
	NoColor bool
	// -i
	// --ignore-case
	// Sorting and filtering branches are case insensitive.
	IgnoreCase bool
	// --omit-empty
	// Do not print a newline after formatted refs where the format expands to the empty string.
	OmitEmpty bool
	// --column[=<options>]
	// Display branch listing in columns.
	Column string
	// --no-column
	// Do not display branch listing in columns.
	NoColumn bool
	// -r
	// --remotes
	// List or delete (if used with -d) remote-tracking branches.
	Remotes bool
	// -a
	// --all
	// List both local and remote-tracking branches.
	All bool
	// --list
	// List branches. With --list, the command may be interpreted as branch creation if <pattern> is not provided.
	List bool
	// --show-current
	// Print the name of the current branch.
	ShowCurrent bool
	// -v
	// --verbose
	// Show SHA-1 and commit subject line for each head, along with relationship to upstream branch (if any). If given twice, print the path of the linked worktree (if any) and the name of the upstream branch, as well (if any).
	Verbose bool
	// -q
	// --quiet
	// Be more quiet when creating or deleting a branch, suppressing non-error messages.
	Quiet bool
	// --abbrev=<n>
	// Alter the sha1s minimum display length in the output listing. The default value is 7 and can be overridden by the core.abbrev config option.
	Abbrev uint64
	// --no-abbrev
	// Display the full sha1s in the output listing rather than abbreviating them.
	NoAbbrev bool
	// --track[=(direct|inherit)]
	// When creating a new branch, set up branch.<name>.remote and branch.<name>.merge configuration entries to mark the start-point as 'upstream' from the new branch.
	Track string
	// --no-track
	// Do not set up 'upstream' configuration, even if the branch.autoSetupMerge configuration variable is true.
	NoTrack bool
	// --recurse-submodules
	// THIS OPTION IS EXPERIMENTAL! Causes the current command to recurse into submodules.
	RecurseSubmodules bool
	// -u <upstream>
	// --set-upstream-to=<upstream>
	// Set up <branchname>'s tracking information.
	SetUpstreamTo string
	// --unset-upstream
	// Remove the upstream information for <branchname>. If no branch is specified it defaults to the current branch.
	UnsetUpstream bool
	// --edit-description
	// Open an editor and edit the text to explain what the branch is for.
	EditDescription bool
	// --contains [<commit>]
	// Only list branches which contain the specified commit (HEAD if not specified).
	Contains bool
	// --no-contains [<commit>]
	// Only list branches which don't contain the specified commit (HEAD if not specified).
	NoContains bool
	// --merged [<commit>]
	// Only list branches whose tips are reachable from the specified commit (HEAD if not specified).
	Merged bool
	// --no-merged [<commit>]
	// Only list branches whose tips are not reachable from the specified commit (HEAD if not specified).
	NoMerged bool
	// <branchname>
	// The name of the branch to create or delete. The new branch name must pass all checks defined by git-check-ref-format.
	Branchname string
	// <start-point>
	// The new branch head will point to this commit. It may be given as a branch name, a commit-id, or a tag.
	StartPoint string
	// <oldbranch>
	// The name of an existing branch to rename.
	Oldbranch string
	// <newbranch>
	// The new name for an existing branch. The same restrictions as for <branchname> apply.
	Newbranch string
	// --sort=<key>
	// Sort based on the key given. Prefix - to sort in descending order of the value. You may use the --sort=<key> option multiple times, in which case the last key becomes the primary key.
	Sort string
	// --points-at <object>
	// Only list branches of the given object.
	PointsAt string
	// --format <format>
	// A string that interpolates %(fieldname) from a branch ref being shown and the object it points at. The format is the same as that of git-for-each-ref.
	Format string
	// <pattern>...
	// Shell wildcard patterns to restrict the output to matching branches. If multiple patterns are given, a branch is shown if it matches any of the patterns.
	Pattern []string
}

func BranchCmd(opts *BranchOptions) *exec.Cmd {
	args := []string{"branch"}

	if opts.Delete {
		args = append(args, "--delete")
	}
	if opts.D {
		args = append(args, "-D")
	}
	if opts.CreateReflog {
		args = append(args, "--create-reflog")
	}
	if opts.Force {
		args = append(args, "--force")
	}
	if opts.Move {
		args = append(args, "--move")
	}
	if opts.M {
		args = append(args, "-M")
	}
	if opts.Copy {
		args = append(args, "--copy")
	}
	if opts.C {
		args = append(args, "-C")
	}
	if opts.Color != "" {
		args = append(args, fmt.Sprintf("--color=%s", opts.Color))
	}
	if opts.NoColor {
		args = append(args, "--no-color")
	}
	if opts.IgnoreCase {
		args = append(args, "--ignore-case")
	}
	if opts.OmitEmpty {
		args = append(args, "--omit-empty")
	}
	if opts.Column != "" {
		args = append(args, fmt.Sprintf("--column=%s", opts.Column))
	}
	if opts.NoColumn {
		args = append(args, "--no-column")
	}
	if opts.Remotes {
		args = append(args, "--remotes")
	}
	if opts.All {
		args = append(args, "--all")
	}
	if opts.List {
		args = append(args, "--list")
	}
	if opts.ShowCurrent {
		args = append(args, "--show-current")
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Abbrev > 0 {
		args = append(args, fmt.Sprintf("--abbrev=%d", opts.Abbrev))
	}
	if opts.NoAbbrev {
		args = append(args, "--no-abbrev")
	}
	if opts.Track != "" {
		args = append(args, fmt.Sprintf("--track=%s", opts.Track))
	}
	if opts.NoTrack {
		args = append(args, "--no-track")
	}
	if opts.RecurseSubmodules {
		args = append(args, "--recurse-submodules")
	}
	if opts.SetUpstreamTo != "" {
		args = append(args, fmt.Sprintf("--set-upstream-to=%s", opts.SetUpstreamTo))
	}
	if opts.UnsetUpstream {
		args = append(args, "--unset-upstream")
	}
	if opts.EditDescription {
		args = append(args, "--edit-description")
	}
	if opts.Contains {
		args = append(args, "--contains ")
	}
	if opts.NoContains {
		args = append(args, "--no-contains ")
	}
	if opts.Merged {
		args = append(args, "--merged ")
	}
	if opts.NoMerged {
		args = append(args, "--no-merged ")
	}
	if opts.Sort != "" {
		args = append(args, fmt.Sprintf("--sort=%s", opts.Sort))
	}
	if opts.PointsAt != "" {
		args = append(args, "--points-at")
		args = append(args, opts.PointsAt)
	}
	if opts.Format != "" {
		args = append(args, "--format")
		args = append(args, opts.Format)
	}
	if opts.Branchname != "" {
		args = append(args, opts.Branchname)
	}
	if opts.StartPoint != "" {
		args = append(args, opts.StartPoint)
	}
	if opts.Oldbranch != "" {
		args = append(args, opts.Oldbranch)
	}
	if opts.Newbranch != "" {
		args = append(args, opts.Newbranch)
	}
	if opts.Pattern != nil {
		args = append(args, opts.Pattern...)
	}

	return execGit(opts.CmdDir, args)
}

func Branch(opts *BranchOptions) ([]byte, error) {
	return run(BranchCmd(opts))
}
