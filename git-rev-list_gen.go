// Code generated for gitexec. DO NOT EDIT.
//
// git-rev-list - Lists commit objects in reverse chronological order
//
// Reference: https://git-scm.com/docs/git-rev-list
package gitexec

import (
	"fmt"
	"os/exec"
)

type RevListOptions struct {
	CmdDir string

	// -<number>
	// -n <number>
	// --max-count=<number>
	// Limit the number of commits to output.
	MaxCount uint64
	// --skip=<number>
	// Skip number commits before starting to show the commit output.
	Skip uint64
	// --since=<date>
	// --after=<date>
	// Show commits more recent than a specific date.
	Since string
	// --since-as-filter=<date>
	// Show all commits more recent than a specific date. This visits all commits in the range, rather than stopping at the first commit which is older than a specific date.
	SinceAsFilter string
	// --until=<date>
	// --before=<date>
	// Show commits older than a specific date.
	Until string
	// --max-age=<timestamp>
	// Limit the commits output to ones older than a specific timestamp.
	MaxAge string
	// --min-age=<timestamp>
	// Limit the commits output to ones newer than a specific timestamp.
	MinAge string
	// --author=<pattern>
	// --committer=<pattern>
	// Limit the commits output to ones with committer header lines that match the specified pattern (regular expression).
	Committer string
	// --grep-reflog=<pattern>
	// Limit the commits output to ones with reflog entries that match the specified pattern (regular expression).
	GrepReflog string
	// --grep=<pattern>
	// Limit the commits output to ones with a log message that matches the specified pattern (regular expression).
	Grep string
	// --all-match
	// Limit the commits output to ones that match all given --grep, instead of ones that match at least one.
	AllMatch bool
	// --invert-grep
	// Limit the commits output to ones with a log message that do not match the pattern specified with --grep=<pattern>.
	InvertGrep bool
	// -i
	// --regexp-ignore-case
	// Match the regular expression limiting patterns without regard to letter case.
	RegexpIgnoreCase bool
	// --basic-regexp
	// Consider the limiting patterns to be basic regular expressions; this is the default.
	BasicRegexp bool
	// -E
	// --extended-regexp
	// Consider the limiting patterns to be extended regular expressions instead of the default basic regular expressions.
	ExtendedRegexp bool
	// -F
	// --fixed-strings
	// Consider the limiting patterns to be fixed strings (don’t interpret pattern as a regular expression).
	FixedStrings bool
	// -P
	// --perl-regexp
	// Consider the limiting patterns to be Perl-compatible regular expressions.
	PerlRegexp bool
	// --remove-empty
	// Stop when a given path disappears from the tree.
	RemoveEmpty bool
	// --merges
	// Print only merge commits. This is exactly the same as --min-parents=2.
	Merges bool
	// --no-merges
	// Do not print commits with more than one parent. This is exactly the same as --max-parents=1.
	NoMerges bool
	// --min-parents=<number>
	// Show only commits which have at least that many parent commits.
	MinParents uint64
	// --max-parents=<number>
	// Show only commits which have at most that many parent commits.
	MaxParents uint64
	// --no-min-parents
	// Resets the min-parents limit.
	NoMinParents bool
	// --no-max-parents
	// Resets the max-parents limit.
	NoMaxParents bool
	// --first-parent
	// Follow only the first parent commit upon seeing a merge commit.
	FirstParent bool
	// --exclude-first-parent-only
	// When finding commits to exclude (with a ^), follow only the first parent commit upon seeing a merge commit.
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
	// --quiet
	// Don't print anything to standard output.
	Quiet bool
	// --disk-usage
	// --disk-usage=human
	// Suppress normal output; instead, print the sum of the bytes used for on-disk storage by the selected commits or objects.
	DiskUsage bool
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
	// --use-bitmap-index
	// Try to speed up the traversal using the pack bitmap index (if one is available).
	UseBitmapIndex bool
	// --progress=<header>
	// Show progress reports on stderr as objects are considered.
	Progress string
	// --simplify-by-decoration
	// Commits that are referred by some branch or tag are selected.
	SimplifyByDecoration bool
	// --show-pulls
	// Include all commits from the default mode, but also any merge commits that are not TREESAME to the first parent but are TREESAME to a later parent.
	ShowPulls bool
	// --full-history
	// Same as the default mode, but does not prune some history.
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
	// When given a range of commits to display (e.g. commit1..commit2 or commit2 ^commit1), and a commit <commit> in that range, only display commits in that range that are ancestors of <commit>, descendants of <commit>, or <commit> itself.
	AncestryPath string
	// --bisect
	// Limit output to the one commit object which is roughly halfway between included and excluded commits.
	Bisect bool
	// --bisect-vars
	// This calculates the same as --bisect, except that refs in refs/bisect/ are not used, and except that this outputs text ready to be eval'ed by the shell.
	BisectVars bool
	// --bisect-all
	// This outputs all the commit objects between the included and excluded commits, ordered by their distance to the included and excluded commits.
	BisectAll bool
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
	// --objects
	// Print the object IDs of any object referenced by the listed commits.
	Objects bool
	// --in-commit-order
	// Print tree and blob ids in order of the commits.
	InCommitOrder bool
	// --objects-edge
	// Similar to --objects, but also print the IDs of excluded commits prefixed with a "-" character.
	ObjectsEdge bool
	// --objects-edge-aggressive
	// Similar to --objects-edge, but it tries harder to find excluded commits at the cost of increased time.
	ObjectsEdgeAggressive bool
	// --indexed-objects
	// Pretend as if all trees and blobs used by the index are listed on the command line.
	IndexedObjects bool
	// --unpacked
	// Only useful with --objects; print the object IDs that are not in packs.
	Unpacked bool
	// --object-names
	// Only useful with --objects; print the names of the object IDs that are found.
	ObjectNames bool
	// --no-object-names
	// Only useful with --objects; does not print the names of the object IDs that are found.
	NoObjectNames bool
	// --filter=<filter-spec>
	// Only useful with one of the --objects*; omits objects (usually blobs) from the list of printed objects.
	Filter string
	// --no-filter
	// Turn off any previous --filter= argument.
	NoFilter bool
	// --filter-provided-objects
	// Filter the list of explicitly provided objects, which would otherwise always be printed even if they did not match any of the filters. Only useful with --filter=.
	FilterProvidedObjects bool
	// --filter-print-omitted
	// Only useful with --filter=; prints a list of the objects omitted by the filter.
	FilterPrintOmitted bool
	// --missing=<missing-action>
	// A debug option to help with future "partial clone" development. This option specifies how missing objects are handled.
	Missing string
	// --exclude-promisor-objects
	// (For internal use only.) Prefilter object traversal at promisor boundary. This is used with partial clone.
	ExcludePromisorObjects bool
	// --no-walk[=(sorted|unsorted)]
	// Only show the given commits, but do not traverse their ancestors.
	NoWalk string
	// --do-walk
	// Overrides a previous --no-walk.
	DoWalk bool
	// --pretty[=<format>]
	// Pretty-print the contents of the commit logs in a given format, where <format> can be one of oneline, short, medium, full, fuller, reference, email, raw, format:<string> and tformat:<string>.
	Pretty string
	// --format=<format>
	// Pretty-print the contents of the commit logs in a given format, where <format> can be one of oneline, short, medium, full, fuller, reference, email, raw, format:<string> and tformat:<string>.
	Format string
	// --abbrev-commit
	// Instead of showing the full 40-byte hexadecimal commit object name, show a prefix that names the object uniquely.
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
	// --show-signature
	// Check the validity of a signed commit object by passing the signature to gpg --verify and show the output.
	ShowSignature bool
	// --relative-date
	// Synonym for --date=relative.
	RelativeDate bool
	// --date=<format>
	// Only takes effect for dates shown in human-readable format, such as when using --pretty.
	Date string
	// --header
	// Print the contents of the commit in raw-format; each record is separated with a NUL character.
	Header bool
	// --no-commit-header
	// Suppress the header line containing "commit" and the object ID printed before the specified format.
	NoCommitHeader bool
	// --commit-header
	// Overrides a previous --no-commit-header.
	CommitHeader bool
	// --parents
	// Print also the parents of the commit (in the form "commit parent..."). Also enables parent rewriting.
	Parents bool
	// --children
	// Print also the children of the commit (in the form "commit parent..."). Also enables parent rewriting.
	Children bool
	// --timestamp
	// Print the raw commit timestamp.
	Timestamp bool
	// --left-right
	// Mark which side of a symmetric difference a commit is reachable from.
	LeftRight bool
	// --graph
	// Draw a text-based graphical representation of the commit history on the left hand side of the output.
	Graph bool
	// --show-linear-break[=<barrier>]
	// When --graph is not used, all history branches are flattened which can make it hard to see that the two consecutive commits do not belong to a linear branch.
	ShowLinearBreak string
	// --count
	// Print a number stating how many commits would have been listed, and suppress all other output.
	Count bool
	// <commit>
	// The commit to compare against. If not specified, the working tree is compared against HEAD.
	Commit string
	// --
	// Do not interpret any more arguments as options.
	DoNotInterpretMoreArgumentsAsOptions bool
	// <path>...
	// Commits modifying the given <paths> are selected.
	Path []string
}

func RevListCmd(opts *RevListOptions) *exec.Cmd {
	args := []string{"rev-list"}

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
	if opts.MaxAge != "" {
		args = append(args, fmt.Sprintf("--max-age=%s", opts.MaxAge))
	}
	if opts.MinAge != "" {
		args = append(args, fmt.Sprintf("--min-age=%s", opts.MinAge))
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
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.DiskUsage {
		args = append(args, "--disk-usage")
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
	if opts.UseBitmapIndex {
		args = append(args, "--use-bitmap-index")
	}
	if opts.Progress != "" {
		args = append(args, fmt.Sprintf("--progress=%s", opts.Progress))
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
	if opts.Bisect {
		args = append(args, "--bisect")
	}
	if opts.BisectVars {
		args = append(args, "--bisect-vars")
	}
	if opts.BisectAll {
		args = append(args, "--bisect-all")
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
	if opts.Objects {
		args = append(args, "--objects")
	}
	if opts.InCommitOrder {
		args = append(args, "--in-commit-order")
	}
	if opts.ObjectsEdge {
		args = append(args, "--objects-edge")
	}
	if opts.ObjectsEdgeAggressive {
		args = append(args, "--objects-edge-aggressive")
	}
	if opts.IndexedObjects {
		args = append(args, "--indexed-objects")
	}
	if opts.Unpacked {
		args = append(args, "--unpacked")
	}
	if opts.ObjectNames {
		args = append(args, "--object-names")
	}
	if opts.NoObjectNames {
		args = append(args, "--no-object-names")
	}
	if opts.Filter != "" {
		args = append(args, fmt.Sprintf("--filter=%s", opts.Filter))
	}
	if opts.NoFilter {
		args = append(args, "--no-filter")
	}
	if opts.FilterProvidedObjects {
		args = append(args, "--filter-provided-objects")
	}
	if opts.FilterPrintOmitted {
		args = append(args, "--filter-print-omitted")
	}
	if opts.Missing != "" {
		args = append(args, fmt.Sprintf("--missing=%s", opts.Missing))
	}
	if opts.ExcludePromisorObjects {
		args = append(args, "--exclude-promisor-objects")
	}
	if opts.NoWalk != "" {
		args = append(args, fmt.Sprintf("--no-walk=%s", opts.NoWalk))
	}
	if opts.DoWalk {
		args = append(args, "--do-walk")
	}
	if opts.Pretty != "" {
		args = append(args, fmt.Sprintf("--pretty=%s", opts.Pretty))
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
	if opts.ShowSignature {
		args = append(args, "--show-signature")
	}
	if opts.RelativeDate {
		args = append(args, "--relative-date")
	}
	if opts.Date != "" {
		args = append(args, fmt.Sprintf("--date=%s", opts.Date))
	}
	if opts.Header {
		args = append(args, "--header")
	}
	if opts.NoCommitHeader {
		args = append(args, "--no-commit-header")
	}
	if opts.CommitHeader {
		args = append(args, "--commit-header")
	}
	if opts.Parents {
		args = append(args, "--parents")
	}
	if opts.Children {
		args = append(args, "--children")
	}
	if opts.Timestamp {
		args = append(args, "--timestamp")
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
	if opts.Count {
		args = append(args, "--count")
	}
	if opts.Commit != "" {
		args = append(args, opts.Commit)
	}
	if opts.DoNotInterpretMoreArgumentsAsOptions {
		args = append(args, "--")
	}
	if opts.Path != nil {
		args = append(args, opts.Path...)
	}

	return execGit(opts.CmdDir, args)
}

func RevList(opts *RevListOptions) ([]byte, error) {
	return run(RevListCmd(opts))
}
