// Code generated for gitexec. DO NOT EDIT.
//
// git-annotate - Annotate file lines with commit information
//
// Reference: https://git-scm.com/docs/git-annotate
package gitexec

import (
	"fmt"
	"os/exec"
)

type AnnotateOptions struct {
	CmdDir string

	// -b
	// Show blank SHA-1 for boundary commits. This can also be controlled via the blame.blankBoundary config option.
	B bool
	// --root
	// Do not treat root commits as boundaries. This can also be controlled via the blame.showRoot config option.
	Root bool
	// --show-stats
	// Include additional statistics at the end of the output.
	ShowStats bool
	// -L <start>
	// <end>
	// -L :<funcname>
	// Annotate only the line range given by <start>,<end>, or by the function name regex <funcname>. May be specified multiple times. Overlapping ranges are allowed.
	Lstartend string
	// -l
	// Show long revision (default: off).
	L bool
	// -t
	// Show raw timestamp (default: off).
	T bool
	// -S <revs-file>
	// Use revisions from revs-file instead of calling git-rev-list.
	Srevsfile string
	// --reverse <rev>..<rev>
	// Walk history forward instead of backward, showing the last revision in which a line has existed. Requires a range of revisions where the path exists in the start revision.
	Reverse string
	// --first-parent
	// Follow only the first parent commit upon seeing a merge commit. Useful for determining when a line was introduced to a particular integration branch.
	FirstParent bool
	// -p
	// --porcelain
	// Show in a format designed for machine consumption.
	Porcelain bool
	// --line-porcelain
	// Show the porcelain format with commit information for each line, not just the first time a commit is referenced. Implies --porcelain.
	LinePorcelain bool
	// --incremental
	// Show the result incrementally in a format designed for machine consumption.
	Incremental bool
	// --encoding=<encoding>
	// Specifies the encoding used to output author names and commit summaries. Set to 'none' for unconverted data.
	Encoding string
	// --contents <file>
	// Annotate using the contents from the named file, starting from <rev> if specified, and HEAD otherwise. Use - to read from standard input.
	Contents string
	// --date=<format>
	// Specifies the format used to output dates. If not provided, uses blame.date config or defaults to 'iso' format.
	Date string
	// --progress
	// Show progress status on stderr. Enabled by default when attached to a terminal. Can't be used with --porcelain or --incremental.
	Progress bool
	// --no-progress
	// Do not show progress status on stderr.
	NoProgress bool
	// -M[<num>]
	// Detect moved or copied lines within a file. Optional <num> is the lower bound of alphanumeric characters that must be detected as moving/copying (default: 20).
	Mnum int
	// -C[<num>]
	// In addition to -M, detect lines moved or copied from other files modified in the same commit. Optional <num> is the lower bound of alphanumeric characters (default: 40).
	Cnum int
	// --ignore-rev <rev>
	// Ignore changes made by the specified revision when assigning blame. May be specified multiple times.
	IgnoreRev string
	// --ignore-revs-file <file>
	// Ignore revisions listed in the specified file, which must be in the same format as fsck.skipList. Use "" to clear the list.
	IgnoreRevsFile string
	// --color-lines
	// Color line annotations differently if they come from the same commit as the preceding line. Color defaults to cyan and can be adjusted via color.blame.repeatedLines.
	ColorLines bool
	// --color-by-age
	// Color line annotations based on the line's age. Colors are controlled by color.blame.highlightRecent.
	ColorByAge bool
	// -h
	// Show help message.
	H bool
	// --
	// Do not interpret any more arguments as options.
	DoNotInterpretMoreArgumentsAsOptions bool
	// <file>
	// The file on which to operate.
	File string
}

func AnnotateCmd(opts *AnnotateOptions) *exec.Cmd {
	args := []string{"annotate"}

	if opts.B {
		args = append(args, "-b")
	}
	if opts.Root {
		args = append(args, "--root")
	}
	if opts.ShowStats {
		args = append(args, "--show-stats")
	}
	if opts.Lstartend != "" {
		args = append(args, "-L")
		args = append(args, opts.Lstartend)
	}
	if opts.L {
		args = append(args, "-l")
	}
	if opts.T {
		args = append(args, "-t")
	}
	if opts.Srevsfile != "" {
		args = append(args, "-S")
		args = append(args, opts.Srevsfile)
	}
	if opts.Reverse != "" {
		args = append(args, "--reverse")
		args = append(args, opts.Reverse)
	}
	if opts.FirstParent {
		args = append(args, "--first-parent")
	}
	if opts.Porcelain {
		args = append(args, "--porcelain")
	}
	if opts.LinePorcelain {
		args = append(args, "--line-porcelain")
	}
	if opts.Incremental {
		args = append(args, "--incremental")
	}
	if opts.Encoding != "" {
		args = append(args, fmt.Sprintf("--encoding=%s", opts.Encoding))
	}
	if opts.Contents != "" {
		args = append(args, "--contents")
		args = append(args, opts.Contents)
	}
	if opts.Date != "" {
		args = append(args, fmt.Sprintf("--date=%s", opts.Date))
	}
	if opts.Progress {
		args = append(args, "--progress")
	}
	if opts.NoProgress {
		args = append(args, "--no-progress")
	}
	if opts.Mnum > 0 {
		args = append(args, fmt.Sprintf("-M%d", opts.Mnum))
	}
	if opts.Cnum > 0 {
		args = append(args, fmt.Sprintf("-C%d", opts.Cnum))
	}
	if opts.IgnoreRev != "" {
		args = append(args, "--ignore-rev")
		args = append(args, opts.IgnoreRev)
	}
	if opts.IgnoreRevsFile != "" {
		args = append(args, "--ignore-revs-file")
		args = append(args, opts.IgnoreRevsFile)
	}
	if opts.ColorLines {
		args = append(args, "--color-lines")
	}
	if opts.ColorByAge {
		args = append(args, "--color-by-age")
	}
	if opts.H {
		args = append(args, "-h")
	}
	if opts.DoNotInterpretMoreArgumentsAsOptions {
		args = append(args, "--")
	}
	if opts.File != "" {
		args = append(args, opts.File)
	}

	return execGit(opts.CmdDir, args)
}

func Annotate(opts *AnnotateOptions) ([]byte, error) {
	return run(AnnotateCmd(opts))
}
