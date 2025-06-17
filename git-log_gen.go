// Code generated for gitexec. DO NOT EDIT.
//
// git-log - Show commit logs
//
// Reference: https://git-scm.com/docs/git-log
package gitexec

import (
	"fmt"
	"os/exec"
)

type LogOptions struct {
	CmdDir string

	// --follow
	// Continue listing the history of a file beyond renames (works only for a single file).
	Follow bool
	// --no-decorate
	// Do not print out ref names of any commits.
	NoDecorate bool
	// --decorate[=short|full|auto|no]
	// Print out ref names of commits. Optionally specify `short`, `full`, `auto`, or `no` to adjust the decoration format (shorthand for `--decorate=short`).
	Decorate string
	// --decorate-refs=<pattern>
	// Include only refs matching the given shell glob pattern in the decorations.
	DecorateRefs string
	// --decorate-refs-exclude=<pattern>
	// Exclude refs matching the given shell glob pattern from the decorations.
	DecorateRefsExclude string
	// --clear-decorations
	// Clear all previous decoration settings and include all refs.
	ClearDecorations bool
	// --source
	// Print out the ref name given on the command line by which each commit was reached.
	Source bool
	// --mailmap
	// Use the mailmap file to canonicalize author and committer names and emails.
	Mailmap bool
	// --no-mailmap
	// Do not use the mailmap file; show author and committer names and emails as recorded.
	NoMailmap bool
	// --use-mailmap
	// Synonym for --mailmap.
	UseMailmap bool
	// --no-use-mailmap
	// Synonym for --no-mailmap.
	NoUseMailmap bool
	// --full-diff
	// Show the full diff for commits touching specified paths, rather than limiting the diff itself.
	FullDiff bool
	// --log-size
	// Include a “log size <number>” line indicating the length (in bytes) of each commit message.
	LogSize bool
	// -L<start>
	// <end>:<file>
	// -L:<funcname>:<file>
	// Trace the evolution of a line range or function within a file. Implies --patch.
	Lstartendfile string
	// <revision-range>
	// Show only commits in the specified revision range. When no <revision-range> is specified, it defaults to HEAD (i.e. the whole history leading to the current commit).
	RevisionRange string
	// --
	// Do not interpret any more arguments as options.
	DoNotInterpretMoreArgumentsAsOptions bool
	// <path>...
	// Show only commits that are enough to explain how the files that match the specified paths came to be.
	Path []string
	// -<number>
	// -n <number>
	// --max-count=<number>
	// Limit the number of commits to output.
	MaxCount int
	// --skip=<number>
	// Skip <number> commits before starting to show the commit output.
	Skip int
	// --since=<date>
	// --after=<date>
	// Show commits more recent than a specific date.
	Since string
	// --since-as-filter=<date>
	// Show all commits more recent than a specific date by visiting the entire commit range instead of stopping at the first older commit.
	SinceAsFilter string
	// --until=<date>
	// --before=<date>
	// Show commits older than a specific date.
	Until string
	// --author=<pattern>
	// Limit output to commits whose author header lines match the given regular expression.
	Author string
	// --committer=<pattern>
	// Limit output to commits whose committer header lines match the given regular expression.
	Committer string
	// --grep-reflog=<pattern>
	// Limit output to commits whose reflog entries match the given regular expression (requires --walk-reflogs).
	GrepReflog string
	// --grep=<pattern>
	// Limit output to commits whose log message matches the given regular expression.
	Grep string
	// --all-match
	// Require that commits match **all** given --grep patterns rather than any one.
	AllMatch bool
	// --invert-grep
	// Limit output to commits whose log messages **do not** match the given --grep pattern.
	InvertGrep bool
	// -i
	// --regexp-ignore-case
	// Match --grep, --author, etc. patterns case-insensitively.
	RegexpIgnoreCase bool
	// --basic-regexp
	// Treat limiting patterns as basic (POSIX) regular expressions (default).
	BasicRegexp bool
	// -E
	// --extended-regexp
	// Treat limiting patterns as extended regular expressions.
	ExtendedRegexp bool
	// -F
	// --fixed-strings
	// Treat limiting patterns as fixed strings (no regex interpretation).
	FixedStrings bool
	// -P
	// --perl-regexp
	// Treat limiting patterns as Perl-compatible regular expressions.
	PerlRegexp bool
	// --remove-empty
	// Stop traversal when a given path disappears from the tree.
	RemoveEmpty bool
	// --merges
	// Print only merge commits (equivalent to --min-parents=2).
	Merges bool
	// --no-merges
	// Do not print merge commits (equivalent to --max-parents=1).
	NoMerges bool
	// --min-parents=<number>
	// Show only commits with **at least** <number> parents.
	MinParents int
	// --max-parents=<number>
	// Show only commits with **at most** <number> parents.
	MaxParents int
	// --no-min-parents
	// Remove any minimum-parent constraint (no lower limit).
	NoMinParents bool
	// --no-max-parents
	// Remove any maximum-parent constraint (no upper limit).
	NoMaxParents bool
	// --first-parent
	// Follow only the first parent when traversing merge commits.
	FirstParent bool
	// --exclude-first-parent-only
	// When using `^` to exclude commits, follow only the first parent for that exclusion.
	ExcludeFirstParentOnly bool
	// --not
	// Reverses the meaning of the ^ prefix for all following revision specifiers.
	Not bool
	// --all
	// Pretend as if all the refs in refs/, along with HEAD, are listed on the command line as <commit>.
	All bool
	// --branches[=<pattern>]
	// Pretend as if all the refs in refs/heads are listed on the command line as <commit>.
	Branches string
	// --tags[=<pattern>]
	// Pretend as if all the refs in refs/tags are listed on the command line as <commit>. If <pattern> is given, limit tags to ones matching given shell glob.
	Tags string
	// --remotes[=<pattern>]
	// Pretend as if all the refs in refs/remotes are listed on the command line as <commit>. If <pattern> is given, limit remote-tracking branches to ones matching given shell glob.
	Remotes string
	// --glob=<glob-pattern>
	// Pretend as if all the refs matching shell glob <glob-pattern> are listed on the command line as <commit>.
	Glob string
	// --exclude=<glob-pattern>
	// Do not include refs matching <glob-pattern> that would otherwise be considered.
	Exclude string
	// --exclude-hidden=[fetch|receive|uploadpack]
	// Do not include refs that would be hidden by git-fetch, git-receive-pack or git-upload-pack by consulting the appropriate fetch.hideRefs, receive.hideRefs or uploadpack.hideRefs configuration along with transfer.hideRefs.
	ExcludeHidden string
	// --reflog
	// Pretend as if all objects mentioned by reflogs are listed on the command line.
	Reflog bool
	// --alternate-refs
	// Pretend as if all objects mentioned as ref tips of alternate repositories were listed on the command line.
	AlternateRefs bool
	// --single-worktree
	// Examine only the current working tree when there are multiple.
	SingleWorktree bool
	// --ignore-missing
	// Upon seeing an invalid object name in the input, pretend as if the bad input was not given.
	IgnoreMissing bool
	// --stdin
	// In addition to getting arguments from the command line, read them from standard input as well.
	Stdin bool
	// --cherry-mark
	// Like --cherry-pick (see below) but mark equivalent commits with = rather than omitting them, and inequivalent ones with +.
	CherryMark bool
	// --cherry-pick
	// Omit any commit that introduces the same change as another commit on the “other side” when the set of commits are limited with symmetric difference.
	CherryPick bool
	// --left-only
	// List only commits on the respective side of a symmetric difference, i.e. only those which would be marked < resp. > by --left-right.
	LeftOnly bool
	// --right-only
	// List only commits on the respective side of a symmetric difference, i.e. only those which would be marked < resp. > by --left-right.
	RightOnly bool
	// --cherry
	// A synonym for --right-only --cherry-mark --no-merges; useful to limit the output to the commits on our side and mark those that have been applied to the other side of a forked history with git log --cherry upstream...mybranch, similar to git cherry upstream mybranch.
	Cherry bool
	// -g
	// --walk-reflogs
	// Instead of walking the commit ancestry chain, walk reflog entries from the most recent one to older ones.
	WalkReflogs bool
	// --merge
	// Show commits touching conflicted paths in the range HEAD...<other>, where <other> is the first existing pseudoref in MERGE_HEAD, CHERRY_PICK_HEAD, REVERT_HEAD or REBASE_HEAD.
	Merge bool
	// --boundary
	// Output excluded boundary commits. Boundary commits are prefixed with -.
	Boundary bool
	// <paths>
	// Commits modifying the given <paths> are selected.
	Paths string
	// --simplify-by-decoration
	// Simplify history by decoration, keeping only commits referenced by tags or affecting the given paths.
	SimplifyByDecoration bool
	// --show-pulls
	// Show pull merges that bring in topic branches (even if simplified away).
	ShowPulls bool
	// --full-history
	// Do not simplify history—show every commit (with parent rewriting).
	FullHistory bool
	// --dense
	// Only the selected commits are shown, plus some to have a meaningful history.
	Dense bool
	// --sparse
	// All commits in the simplified history are shown.
	Sparse bool
	// --simplify-merges
	// Additional option to --full-history to remove some needless merges from the resulting history, as there are no selected commits contributing to this merge.
	SimplifyMerges bool
	// --ancestry-path[=<commit>]
	// Show only commits on the ancestry chain to or from the given <commit>.
	AncestryPath string
	// --date-order
	// Show no parents before all of its children are shown, but otherwise show commits in the commit timestamp order.
	DateOrder bool
	// --author-date-order
	// Show no parents before all of its children are shown, but otherwise show commits in the author timestamp order.
	AuthorDateOrder bool
	// --topo-order
	// Show no parents before all of its children are shown, and avoid showing commits on multiple lines of history intermixed.
	TopoOrder bool
	// --reverse
	// Output the commits chosen to be shown (see Commit Limiting section above) in reverse order.
	Reverse bool
	// --no-walk[=(sorted|unsorted)]
	// Show only the commits given on the command line, without traversing ancestors; unsorted shows in input order, sorted (or no arg) shows in reverse chronological order.
	NoWalk string
	// --do-walk
	// Override a previous --no-walk and resume normal ancestry traversal.
	DoWalk bool
	// --pretty[=<format>]
	// --format=<format>
	// Pretty-print commits using formats such as oneline, short, medium, full, fuller, reference, email, raw, format:<string> or tformat:<string>.
	Format string
	// --abbrev-commit
	// Show abbreviated commit object names instead of full 40-byte hashes.
	AbbrevCommit bool
	// --no-abbrev-commit
	// Show the full 40-byte hexadecimal commit object name.
	NoAbbrevCommit bool
	// --oneline
	// This is a shorthand for "--pretty=oneline --abbrev-commit" used together.
	Oneline bool
	// --encoding=<encoding>
	// This is a shorthand for "--pretty=oneline --abbrev-commit" used together.
	Encoding string
	// --expand-tabs=<n>
	// Perform a tab expansion (replace each tab with enough spaces to fill to the next display column that is a multiple of <n>) in the log message before showing it in the output.
	ExpandTabs uint64
	// --no-expand-tabs
	// Short-hand for --expand-tabs=0, which disables tab expansion.
	NoExpandTabs bool
	// --notes[=<ref>]
	// Show notes from the specified notes ref, or from the default if none is given.
	Notes string
	// --no-notes
	// Do not show any notes, resetting the list of visible notes refs.
	NoNotes bool
	// --show-notes-by-default
	// Show default notes unless specific --notes options are given.
	ShowNotesByDefault bool
	// --show-signature
	// Verify a signed commit object by passing its signature to gpg --verify and show the result.
	ShowSignature bool
	// --relative-date
	// Synonym for --date=relative.
	RelativeDate bool
	// --date=<format>
	// Control date output format (relative, iso, iso-strict, rfc, short, raw, human, unix, format:<string>, format-local:<string>, default).
	Date string
	// --parents
	// Also show parent hashes of each commit (enables parent rewriting).
	Parents bool
	// --children
	// Also show child hashes of each commit (enables parent rewriting).
	Children bool
	// --left-right
	// Prefix commits with `<` or `>` to show which side of a symmetric difference they come from.
	LeftRight bool
	// --graph
	// Draw an ASCII graph of the commit history on the left side of the output.
	Graph bool
	// --show-linear-break[=<barrier>]
	// When --graph is not used, all history branches are flattened which can make it hard to see that the two consecutive commits do not belong to a linear branch.
	ShowLinearBreak string
}

func LogCmd(opts *LogOptions) *exec.Cmd {
	args := []string{"log"}

	if opts.Follow {
		args = append(args, "--follow")
	}
	if opts.NoDecorate {
		args = append(args, "--no-decorate")
	}
	if opts.Decorate != "" {
		args = append(args, fmt.Sprintf("--decorate=%s", opts.Decorate))
	}
	if opts.DecorateRefs != "" {
		args = append(args, fmt.Sprintf("--decorate-refs=%s", opts.DecorateRefs))
	}
	if opts.DecorateRefsExclude != "" {
		args = append(args, fmt.Sprintf("--decorate-refs-exclude=%s", opts.DecorateRefsExclude))
	}
	if opts.ClearDecorations {
		args = append(args, "--clear-decorations")
	}
	if opts.Source {
		args = append(args, "--source")
	}
	if opts.Mailmap {
		args = append(args, "--mailmap")
	}
	if opts.NoMailmap {
		args = append(args, "--no-mailmap")
	}
	if opts.UseMailmap {
		args = append(args, "--use-mailmap")
	}
	if opts.NoUseMailmap {
		args = append(args, "--no-use-mailmap")
	}
	if opts.FullDiff {
		args = append(args, "--full-diff")
	}
	if opts.LogSize {
		args = append(args, "--log-size")
	}
	if opts.Lstartendfile != "" {
		args = append(args, fmt.Sprintf("-L%s", opts.Lstartendfile))
	}
	if opts.RevisionRange != "" {
		args = append(args, opts.RevisionRange)
	}
	if opts.DoNotInterpretMoreArgumentsAsOptions {
		args = append(args, "--")
	}
	if opts.Path != nil {
		args = append(args, opts.Path...)
	}
	if opts.MaxCount > 0 {
		args = append(args, fmt.Sprintf("--max-count=%d", opts.MaxCount))
	}
	if opts.Skip > 0 {
		args = append(args, fmt.Sprintf("--skip=%d", opts.Skip))
	}
	if opts.Since != "" {
		args = append(args, fmt.Sprintf("--since=%s", opts.Since))
	}
	if opts.SinceAsFilter != "" {
		args = append(args, fmt.Sprintf("--since-as-filter=%s", opts.SinceAsFilter))
	}
	if opts.Until != "" {
		args = append(args, fmt.Sprintf("--until=%s", opts.Until))
	}
	if opts.Author != "" {
		args = append(args, fmt.Sprintf("--author=%s", opts.Author))
	}
	if opts.Committer != "" {
		args = append(args, fmt.Sprintf("--committer=%s", opts.Committer))
	}
	if opts.GrepReflog != "" {
		args = append(args, fmt.Sprintf("--grep-reflog=%s", opts.GrepReflog))
	}
	if opts.Grep != "" {
		args = append(args, fmt.Sprintf("--grep=%s", opts.Grep))
	}
	if opts.AllMatch {
		args = append(args, "--all-match")
	}
	if opts.InvertGrep {
		args = append(args, "--invert-grep")
	}
	if opts.RegexpIgnoreCase {
		args = append(args, "--regexp-ignore-case")
	}
	if opts.BasicRegexp {
		args = append(args, "--basic-regexp")
	}
	if opts.ExtendedRegexp {
		args = append(args, "--extended-regexp")
	}
	if opts.FixedStrings {
		args = append(args, "--fixed-strings")
	}
	if opts.PerlRegexp {
		args = append(args, "--perl-regexp")
	}
	if opts.RemoveEmpty {
		args = append(args, "--remove-empty")
	}
	if opts.Merges {
		args = append(args, "--merges")
	}
	if opts.NoMerges {
		args = append(args, "--no-merges")
	}
	if opts.MinParents > 0 {
		args = append(args, fmt.Sprintf("--min-parents=%d", opts.MinParents))
	}
	if opts.MaxParents > 0 {
		args = append(args, fmt.Sprintf("--max-parents=%d", opts.MaxParents))
	}
	if opts.NoMinParents {
		args = append(args, "--no-min-parents")
	}
	if opts.NoMaxParents {
		args = append(args, "--no-max-parents")
	}
	if opts.FirstParent {
		args = append(args, "--first-parent")
	}
	if opts.ExcludeFirstParentOnly {
		args = append(args, "--exclude-first-parent-only")
	}
	if opts.Not {
		args = append(args, "--not")
	}
	if opts.All {
		args = append(args, "--all")
	}
	if opts.Branches != "" {
		args = append(args, fmt.Sprintf("--branches=%s", opts.Branches))
	}
	if opts.Tags != "" {
		args = append(args, fmt.Sprintf("--tags=%s", opts.Tags))
	}
	if opts.Remotes != "" {
		args = append(args, fmt.Sprintf("--remotes=%s", opts.Remotes))
	}
	if opts.Glob != "" {
		args = append(args, fmt.Sprintf("--glob=%s", opts.Glob))
	}
	if opts.Exclude != "" {
		args = append(args, fmt.Sprintf("--exclude=%s", opts.Exclude))
	}
	if opts.ExcludeHidden != "" {
		args = append(args, fmt.Sprintf("--exclude-hidden=%s", opts.ExcludeHidden))
	}
	if opts.Reflog {
		args = append(args, "--reflog")
	}
	if opts.AlternateRefs {
		args = append(args, "--alternate-refs")
	}
	if opts.SingleWorktree {
		args = append(args, "--single-worktree")
	}
	if opts.IgnoreMissing {
		args = append(args, "--ignore-missing")
	}
	if opts.Stdin {
		args = append(args, "--stdin")
	}
	if opts.CherryMark {
		args = append(args, "--cherry-mark")
	}
	if opts.CherryPick {
		args = append(args, "--cherry-pick")
	}
	if opts.LeftOnly {
		args = append(args, "--left-only")
	}
	if opts.RightOnly {
		args = append(args, "--right-only")
	}
	if opts.Cherry {
		args = append(args, "--cherry")
	}
	if opts.WalkReflogs {
		args = append(args, "--walk-reflogs")
	}
	if opts.Merge {
		args = append(args, "--merge")
	}
	if opts.Boundary {
		args = append(args, "--boundary")
	}
	if opts.Paths != "" {
		args = append(args, opts.Paths)
	}
	if opts.SimplifyByDecoration {
		args = append(args, "--simplify-by-decoration")
	}
	if opts.ShowPulls {
		args = append(args, "--show-pulls")
	}
	if opts.FullHistory {
		args = append(args, "--full-history")
	}
	if opts.Dense {
		args = append(args, "--dense")
	}
	if opts.Sparse {
		args = append(args, "--sparse")
	}
	if opts.SimplifyMerges {
		args = append(args, "--simplify-merges")
	}
	if opts.AncestryPath != "" {
		args = append(args, fmt.Sprintf("--ancestry-path=%s", opts.AncestryPath))
	}
	if opts.DateOrder {
		args = append(args, "--date-order")
	}
	if opts.AuthorDateOrder {
		args = append(args, "--author-date-order")
	}
	if opts.TopoOrder {
		args = append(args, "--topo-order")
	}
	if opts.Reverse {
		args = append(args, "--reverse")
	}
	if opts.NoWalk != "" {
		args = append(args, fmt.Sprintf("--no-walk=%s", opts.NoWalk))
	}
	if opts.DoWalk {
		args = append(args, "--do-walk")
	}
	if opts.Format != "" {
		args = append(args, fmt.Sprintf("--format=%s", opts.Format))
	}
	if opts.AbbrevCommit {
		args = append(args, "--abbrev-commit")
	}
	if opts.NoAbbrevCommit {
		args = append(args, "--no-abbrev-commit")
	}
	if opts.Oneline {
		args = append(args, "--oneline")
	}
	if opts.Encoding != "" {
		args = append(args, fmt.Sprintf("--encoding=%s", opts.Encoding))
	}
	if opts.ExpandTabs > 0 {
		args = append(args, fmt.Sprintf("--expand-tabs=%d", opts.ExpandTabs))
	}
	if opts.NoExpandTabs {
		args = append(args, "--no-expand-tabs")
	}
	if opts.Notes != "" {
		args = append(args, fmt.Sprintf("--notes=%s", opts.Notes))
	}
	if opts.NoNotes {
		args = append(args, "--no-notes")
	}
	if opts.ShowNotesByDefault {
		args = append(args, "--show-notes-by-default")
	}
	if opts.ShowSignature {
		args = append(args, "--show-signature")
	}
	if opts.RelativeDate {
		args = append(args, "--relative-date")
	}
	if opts.Date != "" {
		args = append(args, fmt.Sprintf("--date=%s", opts.Date))
	}
	if opts.Parents {
		args = append(args, "--parents")
	}
	if opts.Children {
		args = append(args, "--children")
	}
	if opts.LeftRight {
		args = append(args, "--left-right")
	}
	if opts.Graph {
		args = append(args, "--graph")
	}
	if opts.ShowLinearBreak != "" {
		args = append(args, fmt.Sprintf("--show-linear-break=%s", opts.ShowLinearBreak))
	}

	return execGit(opts.CmdDir, args)
}

func Log(opts *LogOptions) ([]byte, error) {
	return run(LogCmd(opts))
}
