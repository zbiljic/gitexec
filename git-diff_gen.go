// Code generated for gitexec. DO NOT EDIT.
//
// git-diff - Show changes between commits, commit and working tree, etc
//
// Reference: https://git-scm.com/docs/git-diff
package gitexec

import (
	"fmt"
	"os/exec"
)

type DiffOptions struct {
	CmdDir string

	// -p
	// -u
	// --patch
	// Generate patch (see Generating patch text with -p). This is the default.
	Patch bool
	// -s
	// --no-patch
	// Suppress all output from the diff machinery. Useful for commands like git show that show the patch by default to squelch their output, or to cancel the effect of options like --patch, --stat earlier on the command line in an alias.
	NoPatch bool
	// -U<n>
	// --unified=<n>
	// Generate diffs with <n> lines of context instead of the usual three. Implies --patch.
	Unified uint64
	// --output=<file>
	// Output to a specific file instead of stdout.
	Output string
	// --output-indicator-new=<char>
	// Specify the character used to indicate new lines in the generated patch. Normally it is +.
	OutputIndicatorNew string
	// --output-indicator-old=<char>
	// Specify the character used to indicate old lines in the generated patch. Normally it is -.
	OutputIndicatorOld string
	// --output-indicator-context=<char>
	// Specify the character used to indicate context lines in the generated patch. Normally it is ' '.
	OutputIndicatorContext string
	// --raw
	// Generate the diff in raw format.
	Raw bool
	// --patch-with-raw
	// Synonym for -p --raw.
	PatchWithRaw bool
	// --indent-heuristic
	// Enable the heuristic that shifts diff hunk boundaries to make patches easier to read. This is the default.
	IndentHeuristic bool
	// --no-indent-heuristic
	// Disable the indent heuristic.
	NoIndentHeuristic bool
	// --minimal
	// Spend extra time to make sure the smallest possible diff is produced.
	Minimal bool
	// --patience
	// Generate a diff using the "patience diff" algorithm.
	Patience bool
	// --histogram
	// Generate a diff using the "histogram diff" algorithm.
	Histogram bool
	// --anchored=<text>
	// Generate a diff using the "anchored diff" algorithm. This option may be specified more than once.
	Anchored string
	// --diff-algorithm=(patience|minimal|histogram|myers|default)
	// Choose a diff algorithm. Variants: myers (default), minimal, patience, histogram.
	DiffAlgorithm string
	// --stat[=<width>[
	// <name-width>[
	// <count>]]]
	// Generate a diffstat. Control width, name-width, and count.
	Stat string
	// --compact-summary
	// Output a condensed summary of extended header information in diffstat. Implies --stat.
	CompactSummary bool
	// --numstat
	// Similar to --stat, but shows number of added and deleted lines in decimal notation and pathname without abbreviation.
	Numstat bool
	// --shortstat
	// Output only the last line of the --stat format containing total number of modified files, as well as number of added and deleted lines.
	Shortstat bool
	// -X [<param1>
	// <param2>
	// ...]
	// --dirstat[=<param1>
	// <param2>
	// ...]
	// Output the distribution of relative amount of changes for each sub-directory. Parameters include: changes, lines.
	Dirstat string
	// --cumulative
	// Synonym for --dirstat=cumulative.
	Cumulative bool
	// --dirstat-by-file[=<param>
	// ...]
	// Synonym for --dirstat=files,<param>,... .
	DirstatByFile string
	// --summary
	// Output a condensed summary of extended header information such as creations, renames and mode changes.
	Summary bool
	// --patch-with-stat
	// Synonym for -p --stat.
	PatchWithStat bool
	// -z
	// When --raw, --numstat, --name-only or --name-status has been given, do not munge pathnames and use NULs as output field terminators.
	Z bool
	// --name-only
	// Show only the name of each changed file in the post-image tree.
	NameOnly bool
	// --name-status
	// Show only the name(s) and status of each changed file.
	NameStatus bool
	// --submodule[=<format>]
	// Specify how differences in submodules are shown. When specifying --submodule=short the short format is used.
	Submodule string
	// --color[=<when>]
	// Show colored diff. --color (i.e. without =<when>) is the same as --color=always. <when> can be one of always, never, or auto.
	Color string
	// --no-color
	// Turn off colored diff.
	NoColor bool
	// --color-moved[=<mode>]
	// Moved lines of code are colored differently.
	ColorMoved string
	// --no-color-moved
	// Turn off move detection.
	NoColorMoved bool
	// --color-moved-ws=<mode>
	// ...
	// This configures how whitespace is ignored when performing the move detection for --color-moved.
	ColorMovedWs string
	// --no-color-moved-ws
	// Do not ignore whitespace when performing move detection.
	NoColorMovedWs bool
	// --word-diff[=<mode>]
	// By default, words are delimited by whitespace; see --word-diff-regex below.
	WordDiff string
	// --word-diff-regex=<regex>
	// Use <regex> to decide what a word is, instead of considering runs of non-whitespace to be a word. Also implies --word-diff unless it was already enabled.
	WordDiffRegex string
	// --color-words[=<regex>]
	// Equivalent to --word-diff=color plus (if a regex was specified) --word-diff-regex=<regex>.
	ColorWords string
	// --no-renames
	// Turn off rename detection, even when the configuration file gives the default to do so.
	NoRenames bool
	// --rename-empty
	// Whether to use empty blobs as rename source.
	RenameEmpty bool
	// --no-rename-empty
	// Prevent use of empty blobs as rename source.
	NoRenameEmpty bool
	// --check
	// Warn if changes introduce conflict markers or whitespace errors.
	Check bool
	// --ws-error-highlight=<kind>
	// Highlight whitespace errors in the context, old or new lines of the diff. Multiple values are separated by comma.
	WsErrorHighlight string
	// --full-index
	// Instead of the first handful of characters, show the full pre- and post-image blob object names on the "index" line when generating patch format output.
	FullIndex bool
	// --binary
	// In addition to --full-index, output a binary diff that can be applied with git-apply. Implies --patch.
	Binary bool
	// --abbrev[=<n>]
	// Instead of showing the full 40-byte hexadecimal object name in diff-raw format output and diff-tree header lines, show the shortest prefix that is at least <n> hexdigits long that uniquely refers the object.
	Abbrev uint64
	// -B[<n>][/<m>]
	// --break-rewrites[=[<n>][/<m>]]
	// Break complete rewrite changes into pairs of delete and create.
	BreakRewrites uint64
	// -M[<n>]
	// --find-renames[=<n>]
	// Detect renames. If <n> is specified, it is a threshold on the similarity index (i.e. amount of addition/deletions compared to the file's size).
	FindRenames uint64
	// -C[<n>]
	// --find-copies[=<n>]
	// Detect copies as well as renames.
	FindCopies uint64
	// --find-copies-harder
	// For performance reasons, by default, -C option finds copies only if the original file of the copy was modified in the same changeset. This flag makes the command inspect unmodified files as candidates for the source of copy.
	FindCopiesHarder bool
	// -D
	// --irreversible-delete
	// .
	IrreversibleDelete bool
	// -l<num>
	// This option prevents the exhaustive portion of rename/copy detection from running if the number of source/destination files involved exceeds the specified number.
	L uint64
	// --diff-filter=[(A|C|D|M|R|T|U|X|B)...[*]]
	// Select only files that are Added (A), Copied (C), Deleted (D), Modified (M), Renamed (R), have their type (i.e. regular file, symlink, submodule, …​) changed (T), are Unmerged (U), are Unknown (X), or have had their pairing Broken (B). Any combination of the filter characters (including none) can be used.
	DiffFilter string
	// -S<string>
	// Look for differences that change the number of occurrences of the specified <string> (i.e. addition/deletion) in a file. Intended for the scripter's use.
	S bool
	// -G<regex>
	// Look for differences whose patch text contains added/removed lines that match <regex>.
	G bool
	// --find-object=<object-id>
	// Look for differences that change the number of occurrences of the specified object. Similar to -S, just the argument is different in that it doesn't search for a specific string but for a specific object id.
	FindObject string
	// --pickaxe-all
	// When -S or -G finds a change, show all the changes in that changeset, not just the files that contain the change in <string>.
	PickaxeAll bool
	// --pickaxe-regex
	// Treat the <string> given to -S as an extended POSIX regular expression to match.
	PickaxeRegex bool
	// -O<orderfile>
	// Control the order in which files appear in the output.
	O bool
	// --skip-to=<file>
	// Discard the files before the named <file> from the output (i.e. skip to).
	SkipTo string
	// --rotate-to=<file>
	// Move the files before the named <file> to the end of the output (i.e. rotate to).
	RotateTo string
	// -R
	// Swap two inputs; that is, show differences from index or on-disk file to tree contents.
	R bool
	// --relative[=<path>]
	// When run from a subdirectory of the project, it can be told to exclude changes outside the directory and show pathnames relative to it with this option.
	Relative string
	// --no-relative
	// Used to countermand both diff.relative config option and previous --relative.
	NoRelative bool
	// -a
	// --text
	// Treat all files as text.
	Text bool
	// --ignore-cr-at-eol
	// Ignore carriage-return at the end of line when doing a comparison.
	IgnoreCrAtEol bool
	// --ignore-space-at-eol
	// Ignore changes in whitespace at EOL.
	IgnoreSpaceAtEol bool
	// -b
	// --ignore-space-change
	// Ignore changes in amount of whitespace. This ignores whitespace at line end, and considers all other sequences of one or more whitespace characters to be equivalent.
	IgnoreSpaceChange bool
	// -w
	// --ignore-all-space
	// Ignore whitespace when comparing lines. This ignores differences even if one line has whitespace where the other line has none.
	IgnoreAllSpace bool
	// --ignore-blank-lines
	// Ignore changes whose lines are all blank.
	IgnoreBlankLines bool
	// -I<regex>
	// --ignore-matching-lines=<regex>
	// Ignore changes whose all lines match <regex>. This option may be specified more than once.
	IgnoreMatchingLines string
	// --inter-hunk-context=<number>
	// Show the context between diff hunks, up to the specified <number> of lines, thereby fusing hunks that are close to each other.
	InterHunkContext string
	// -W
	// --function-context
	// Show whole function as context lines for each change.
	FunctionContext bool
	// --exit-code
	// Make the program exit with codes similar to diff. That is, it exits with 1 if there were differences and 0 means no differences.
	ExitCode bool
	// --quiet
	// Disable all output of the program. Implies --exit-code.
	Quiet bool
	// --ext-diff
	// Allow an external diff helper to be executed.
	ExtDiff bool
	// --no-ext-diff
	// Disallow external diff drivers.
	NoExtDiff bool
	// --textconv
	// Allow external text conversion filters to be run when comparing binary files.
	Textconv bool
	// --no-textconv
	// Disallow external text conversion filters to be run when comparing binary files.
	NoTextconv bool
	// --ignore-submodules[=(none|untracked|dirty|all)]
	// Ignore changes to submodules in the diff generation. all is the default.
	IgnoreSubmodules string
	// --src-prefix=<prefix>
	// Show the given source <prefix> instead of "a/".
	SrcPrefix string
	// --dst-prefix=<prefix>
	// Show the given destination <prefix> instead of "b/".
	DstPrefix string
	// --no-prefix
	// Do not show any source or destination prefix.
	NoPrefix bool
	// --default-prefix
	// Use the default source and destination prefixes ("a/" and "b/").
	DefaultPrefix bool
	// --line-prefix=<prefix>
	// Prepend an additional <prefix> to every line of output.
	LinePrefix string
	// --ita-invisible-in-index
	// This option makes the entry appear as a new file in git diff and non-existent in git diff --cached.
	ItaInvisibleInIndex bool
	// -1
	// --base
	// Compare the working tree with the "base" version (stage #1).
	Base bool
	// -2
	// --ours
	// Compare the working tree with the "our branch" (stage #2).
	Ours bool
	// -3
	// --theirs
	// Compare the working tree with the "their branch" (stage #3).
	Theirs bool
	// -0
	// Omit diff output for unmerged entries and just show "Unmerged". Can be used only when comparing the working tree with the index.
	Zero bool
	// <path>...
	// The <path> parameters, when given, are used to limit the diff to the named paths (you can give directory names and get diff for all files under them).
	Path []string
}

func DiffCmd(opts *DiffOptions) *exec.Cmd {
	args := []string{"diff"}

	if opts.Patch {
		args = append(args, "--patch")
	}
	if opts.NoPatch {
		args = append(args, "--no-patch")
	}
	if opts.Unified > 0 {
		args = append(args, fmt.Sprintf("--unified=%d", opts.Unified))
	}
	if opts.Output != "" {
		args = append(args, fmt.Sprintf("--output=%s", opts.Output))
	}
	if opts.OutputIndicatorNew != "" {
		args = append(args, fmt.Sprintf("--output-indicator-new=%s", opts.OutputIndicatorNew))
	}
	if opts.OutputIndicatorOld != "" {
		args = append(args, fmt.Sprintf("--output-indicator-old=%s", opts.OutputIndicatorOld))
	}
	if opts.OutputIndicatorContext != "" {
		args = append(args, fmt.Sprintf("--output-indicator-context=%s", opts.OutputIndicatorContext))
	}
	if opts.Raw {
		args = append(args, "--raw")
	}
	if opts.PatchWithRaw {
		args = append(args, "--patch-with-raw")
	}
	if opts.IndentHeuristic {
		args = append(args, "--indent-heuristic")
	}
	if opts.NoIndentHeuristic {
		args = append(args, "--no-indent-heuristic")
	}
	if opts.Minimal {
		args = append(args, "--minimal")
	}
	if opts.Patience {
		args = append(args, "--patience")
	}
	if opts.Histogram {
		args = append(args, "--histogram")
	}
	if opts.Anchored != "" {
		args = append(args, fmt.Sprintf("--anchored=%s", opts.Anchored))
	}
	if opts.DiffAlgorithm != "" {
		args = append(args, fmt.Sprintf("--diff-algorithm=%s", opts.DiffAlgorithm))
	}
	if opts.Stat != "" {
		args = append(args, fmt.Sprintf("--stat=%s", opts.Stat))
	}
	if opts.CompactSummary {
		args = append(args, "--compact-summary")
	}
	if opts.Numstat {
		args = append(args, "--numstat")
	}
	if opts.Shortstat {
		args = append(args, "--shortstat")
	}
	if opts.Dirstat != "" {
		args = append(args, fmt.Sprintf("--dirstat=%s", opts.Dirstat))
	}
	if opts.Cumulative {
		args = append(args, "--cumulative")
	}
	if opts.DirstatByFile != "" {
		args = append(args, fmt.Sprintf("--dirstat-by-file=%s", opts.DirstatByFile))
	}
	if opts.Summary {
		args = append(args, "--summary")
	}
	if opts.PatchWithStat {
		args = append(args, "--patch-with-stat")
	}
	if opts.Z {
		args = append(args, "-z")
	}
	if opts.NameOnly {
		args = append(args, "--name-only")
	}
	if opts.NameStatus {
		args = append(args, "--name-status")
	}
	if opts.Submodule != "" {
		args = append(args, fmt.Sprintf("--submodule=%s", opts.Submodule))
	}
	if opts.Color != "" {
		args = append(args, fmt.Sprintf("--color=%s", opts.Color))
	}
	if opts.NoColor {
		args = append(args, "--no-color")
	}
	if opts.ColorMoved != "" {
		args = append(args, fmt.Sprintf("--color-moved=%s", opts.ColorMoved))
	}
	if opts.NoColorMoved {
		args = append(args, "--no-color-moved")
	}
	if opts.ColorMovedWs != "" {
		args = append(args, fmt.Sprintf("--color-moved-ws=%s", opts.ColorMovedWs))
	}
	if opts.NoColorMovedWs {
		args = append(args, "--no-color-moved-ws")
	}
	if opts.WordDiff != "" {
		args = append(args, fmt.Sprintf("--word-diff=%s", opts.WordDiff))
	}
	if opts.WordDiffRegex != "" {
		args = append(args, fmt.Sprintf("--word-diff-regex=%s", opts.WordDiffRegex))
	}
	if opts.ColorWords != "" {
		args = append(args, fmt.Sprintf("--color-words=%s", opts.ColorWords))
	}
	if opts.NoRenames {
		args = append(args, "--no-renames")
	}
	if opts.RenameEmpty {
		args = append(args, "--rename-empty")
	}
	if opts.NoRenameEmpty {
		args = append(args, "--no-rename-empty")
	}
	if opts.Check {
		args = append(args, "--check")
	}
	if opts.WsErrorHighlight != "" {
		args = append(args, fmt.Sprintf("--ws-error-highlight=%s", opts.WsErrorHighlight))
	}
	if opts.FullIndex {
		args = append(args, "--full-index")
	}
	if opts.Binary {
		args = append(args, "--binary")
	}
	if opts.Abbrev > 0 {
		args = append(args, fmt.Sprintf("--abbrev=%d", opts.Abbrev))
	}
	if opts.BreakRewrites > 0 {
		args = append(args, fmt.Sprintf("--break-rewrites=%d", opts.BreakRewrites))
	}
	if opts.FindRenames > 0 {
		args = append(args, fmt.Sprintf("--find-renames=%d", opts.FindRenames))
	}
	if opts.FindCopies > 0 {
		args = append(args, fmt.Sprintf("--find-copies=%d", opts.FindCopies))
	}
	if opts.FindCopiesHarder {
		args = append(args, "--find-copies-harder")
	}
	if opts.IrreversibleDelete {
		args = append(args, "--irreversible-delete")
	}
	if opts.L > 0 {
		args = append(args, fmt.Sprintf("-l%d", opts.L))
	}
	if opts.DiffFilter != "" {
		args = append(args, fmt.Sprintf("--diff-filter=%s", opts.DiffFilter))
	}
	if opts.S {
		args = append(args, "-S<string>")
	}
	if opts.G {
		args = append(args, "-G<regex>")
	}
	if opts.FindObject != "" {
		args = append(args, fmt.Sprintf("--find-object=%s", opts.FindObject))
	}
	if opts.PickaxeAll {
		args = append(args, "--pickaxe-all")
	}
	if opts.PickaxeRegex {
		args = append(args, "--pickaxe-regex")
	}
	if opts.O {
		args = append(args, "-O<orderfile>")
	}
	if opts.SkipTo != "" {
		args = append(args, fmt.Sprintf("--skip-to=%s", opts.SkipTo))
	}
	if opts.RotateTo != "" {
		args = append(args, fmt.Sprintf("--rotate-to=%s", opts.RotateTo))
	}
	if opts.R {
		args = append(args, "-R")
	}
	if opts.Relative != "" {
		args = append(args, fmt.Sprintf("--relative=%s", opts.Relative))
	}
	if opts.NoRelative {
		args = append(args, "--no-relative")
	}
	if opts.Text {
		args = append(args, "--text")
	}
	if opts.IgnoreCrAtEol {
		args = append(args, "--ignore-cr-at-eol")
	}
	if opts.IgnoreSpaceAtEol {
		args = append(args, "--ignore-space-at-eol")
	}
	if opts.IgnoreSpaceChange {
		args = append(args, "--ignore-space-change")
	}
	if opts.IgnoreAllSpace {
		args = append(args, "--ignore-all-space")
	}
	if opts.IgnoreBlankLines {
		args = append(args, "--ignore-blank-lines")
	}
	if opts.IgnoreMatchingLines != "" {
		args = append(args, fmt.Sprintf("--ignore-matching-lines=%s", opts.IgnoreMatchingLines))
	}
	if opts.InterHunkContext != "" {
		args = append(args, fmt.Sprintf("--inter-hunk-context=%s", opts.InterHunkContext))
	}
	if opts.FunctionContext {
		args = append(args, "--function-context")
	}
	if opts.ExitCode {
		args = append(args, "--exit-code")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.ExtDiff {
		args = append(args, "--ext-diff")
	}
	if opts.NoExtDiff {
		args = append(args, "--no-ext-diff")
	}
	if opts.Textconv {
		args = append(args, "--textconv")
	}
	if opts.NoTextconv {
		args = append(args, "--no-textconv")
	}
	if opts.IgnoreSubmodules != "" {
		args = append(args, fmt.Sprintf("--ignore-submodules=%s", opts.IgnoreSubmodules))
	}
	if opts.SrcPrefix != "" {
		args = append(args, fmt.Sprintf("--src-prefix=%s", opts.SrcPrefix))
	}
	if opts.DstPrefix != "" {
		args = append(args, fmt.Sprintf("--dst-prefix=%s", opts.DstPrefix))
	}
	if opts.NoPrefix {
		args = append(args, "--no-prefix")
	}
	if opts.DefaultPrefix {
		args = append(args, "--default-prefix")
	}
	if opts.LinePrefix != "" {
		args = append(args, fmt.Sprintf("--line-prefix=%s", opts.LinePrefix))
	}
	if opts.ItaInvisibleInIndex {
		args = append(args, "--ita-invisible-in-index")
	}
	if opts.Base {
		args = append(args, "--base")
	}
	if opts.Ours {
		args = append(args, "--ours")
	}
	if opts.Theirs {
		args = append(args, "--theirs")
	}
	if opts.Zero {
		args = append(args, "-0")
	}
	if opts.Path != nil {
		args = append(args, opts.Path...)
	}

	return execGit(opts.CmdDir, args)
}

func Diff(opts *DiffOptions) ([]byte, error) {
	return run(DiffCmd(opts))
}
