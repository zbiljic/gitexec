// Code generated for gitexec. DO NOT EDIT.
//
// git-apply - Apply a patch to files and/or to the index
//
// Reference: https://git-scm.com/docs/git-apply
package gitexec

import (
	"fmt"
	"os/exec"
)

type ApplyOptions struct {
	CmdDir string

	// --stat
	// Instead of applying the patch, output diffstat for the input. Turns off "apply".
	Stat bool
	// --numstat
	// Similar to --stat, but shows the number of added and deleted lines in decimal notation and the pathname without abbreviation, to make it more machine friendly.
	Numstat bool
	// --summary
	// Instead of applying the patch, output a condensed summary of information obtained from git diff extended headers, such as creations, renames and mode changes.
	Summary bool
	// --check
	// Instead of applying the patch, see if the patch is applicable to the current working tree and/or the index file and detects errors.
	Check bool
	// --index
	// When --check is in effect, or when applying the patch (which is the default when none of the options that disables it is in effect), make sure the patch is applicable to what the current index file records.
	Index bool
	// --cached
	// Apply a patch without touching the working tree. Instead take the cached data, apply the patch, and store the result in the index without using the working tree.
	Cached bool
	// --intent-to-add
	// When applying the patch only to the working tree, mark new files to be added to the index later.
	IntentToAdd bool
	// -3
	// --3way
	// When the patch does not apply cleanly, fall back on 3-way merge if the patch records the identity of blobs it is supposed to apply to and we have those blobs available locally.
	Threeway bool
	// --ours
	// Instead of leaving conflicts in the file, resolve conflicts favouring our side of the lines. Requires --3way.
	Ours bool
	// --theirs
	// Instead of leaving conflicts in the file, resolve conflicts favouring their side of the lines. Requires --3way.
	Theirs bool
	// --union
	// Instead of leaving conflicts in the file, resolve conflicts favouring both sides of the lines. Requires --3way.
	Union bool
	// --build-fake-ancestor=<file>
	// Newer git diff output has embedded index information for each blob to help identify the original version that the patch applies to.
	BuildFakeAncestor string
	// -R
	// --reverse
	// Apply the patch in reverse.
	Reverse bool
	// --reject
	// For atomicity, git apply by default fails the whole patch and does not touch the working tree when some of the hunks do not apply. This option makes it apply the parts of the patch that are applicable, and leave the rejected hunks in corresponding *.rej files.
	Reject bool
	// -z
	// When --numstat has been given, do not munge pathnames, but use a NUL-terminated machine-readable format.
	Z bool
	// -p<n>
	// Remove <n> leading slashes from traditional diff paths. The default is 1.
	Pn uint64
	// -C<n>
	// Ensure at least <n> lines of surrounding context match before and after each change. When fewer lines of surrounding context exist they all must match.
	Cn uint64
	// --unidiff-zero
	// By default, git apply expects that the patch being applied is a unified diff with at least one line of context. This provides good safety measures, but breaks down when applying a diff generated with --unified=0.
	UnidiffZero bool
	// --apply
	// If you use any of the options marked "Turns off apply" above, git apply reads and outputs the requested information without actually applying the patch. Give this flag after those flags to also apply the patch.
	Apply bool
	// --no-add
	// When applying a patch, ignore additions made by the patch. This can be used to extract the common part between two files by first running diff on them and then applying the result with this option, which would apply the deletion part but not the addition part.
	NoAdd bool
	// --binary
	// --allow-binary-replacement
	// Historically we did not allow binary patch applied without an explicit permission from the user, and this flag was the way to do so. Currently we always allow binary patch application, so this is a no-op.
	AllowBinaryReplacement bool
	// --exclude=<path-pattern>
	// Don't apply changes to files matching the given path pattern. This can be useful when importing patchsets, where you want to exclude certain files or directories.
	Exclude string
	// --include=<path-pattern>
	// Apply changes to files matching the given path pattern. This can be useful when importing patchsets, where you want to include certain files or directories.
	Include string
	// --ignore-space-change
	// --ignore-whitespace
	// When applying a patch, ignore changes in whitespace in context lines if necessary.
	IgnoreWhitespace bool
	// --whitespace=<option>
	// When applying a patch, detect a new or modified line that has whitespace errors. What are considered whitespace errors is controlled by core.whitespace configuration.
	Whitespace string
	// --inaccurate-eof
	// Under certain circumstances, some versions of diff do not correctly detect a missing new-line at the end of the file.
	InaccurateEof bool
	// -v
	// --verbose
	// Report progress to stderr. By default, only a message about the current patch being applied will be printed.
	Verbose bool
	// -q
	// --quiet
	// Suppress stderr output. Messages about patch status and progress will not be printed.
	Quiet bool
	// --recount
	// Do not trust the line counts in the hunk headers, but infer them by inspecting the patch.
	Recount bool
	// --directory=<root>
	// Prepend <root> to all filenames. If a "-p" argument was also passed, it is applied before prepending the new root.
	Directory string
	// --unsafe-paths
	// By default, a patch that affects outside the working area (either a Git controlled working tree, or the current working directory when "git apply" is used as a replacement of GNU patch) is rejected as a mistake.
	UnsafePaths bool
	// --allow-empty
	// Don't return error for patches containing no diff. This includes empty patches and patches with commit text only.
	AllowEmpty bool
	// --allow-overlap
	// Instead of failing when the patch contains changes to submodules, attempt to apply the changes.
	AllowOverlap bool
	// <patch>...
	// The files to read patch from. '-' can be used to read from the standard input.
	Patch []string
}

func ApplyCmd(opts *ApplyOptions) *exec.Cmd {
	args := []string{"apply"}

	if opts.Stat {
		args = append(args, "--stat")
	}
	if opts.Numstat {
		args = append(args, "--numstat")
	}
	if opts.Summary {
		args = append(args, "--summary")
	}
	if opts.Check {
		args = append(args, "--check")
	}
	if opts.Index {
		args = append(args, "--index")
	}
	if opts.Cached {
		args = append(args, "--cached")
	}
	if opts.IntentToAdd {
		args = append(args, "--intent-to-add")
	}
	if opts.Threeway {
		args = append(args, "--3way")
	}
	if opts.Ours {
		args = append(args, "--ours")
	}
	if opts.Theirs {
		args = append(args, "--theirs")
	}
	if opts.Union {
		args = append(args, "--union")
	}
	if opts.BuildFakeAncestor != "" {
		args = append(args, fmt.Sprintf("--build-fake-ancestor=%s", opts.BuildFakeAncestor))
	}
	if opts.Reverse {
		args = append(args, "--reverse")
	}
	if opts.Reject {
		args = append(args, "--reject")
	}
	if opts.Z {
		args = append(args, "-z")
	}
	if opts.Pn > 0 {
		args = append(args, fmt.Sprintf("-p%d", opts.Pn))
	}
	if opts.Cn > 0 {
		args = append(args, fmt.Sprintf("-C%d", opts.Cn))
	}
	if opts.UnidiffZero {
		args = append(args, "--unidiff-zero")
	}
	if opts.Apply {
		args = append(args, "--apply")
	}
	if opts.NoAdd {
		args = append(args, "--no-add")
	}
	if opts.AllowBinaryReplacement {
		args = append(args, "--allow-binary-replacement")
	}
	if opts.Exclude != "" {
		args = append(args, fmt.Sprintf("--exclude=%s", opts.Exclude))
	}
	if opts.Include != "" {
		args = append(args, fmt.Sprintf("--include=%s", opts.Include))
	}
	if opts.IgnoreWhitespace {
		args = append(args, "--ignore-whitespace")
	}
	if opts.Whitespace != "" {
		args = append(args, fmt.Sprintf("--whitespace=%s", opts.Whitespace))
	}
	if opts.InaccurateEof {
		args = append(args, "--inaccurate-eof")
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Recount {
		args = append(args, "--recount")
	}
	if opts.Directory != "" {
		args = append(args, fmt.Sprintf("--directory=%s", opts.Directory))
	}
	if opts.UnsafePaths {
		args = append(args, "--unsafe-paths")
	}
	if opts.AllowEmpty {
		args = append(args, "--allow-empty")
	}
	if opts.AllowOverlap {
		args = append(args, "--allow-overlap")
	}
	if opts.Patch != nil {
		args = append(args, opts.Patch...)
	}

	return execGit(opts.CmdDir, args)
}

func Apply(opts *ApplyOptions) ([]byte, error) {
	return run(ApplyCmd(opts))
}
