// Code generated for gitexec. DO NOT EDIT.
//
// git-commit - Record changes to the repository
//
// Reference: https://git-scm.com/docs/git-commit
package gitexec

import (
	"fmt"
	"os/exec"
)

type CommitOptions struct {
	CmdDir string

	// -a
	// --all
	// Automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.
	All bool
	// -p
	// --patch
	// Use the interactive patch selection interface to choose which changes to commit. See git-add[1] for details.
	Patch bool
	// -C <commit>
	// --reuse-message=<commit>
	// Take an existing <commit> object, and reuse the log message and the authorship information (including the timestamp) when creating the commit.
	ReuseMessage string
	// -c <commit>
	// --reedit-message=<commit>
	// Like -C, but with -c the editor is invoked, so that the user can further edit the commit message.
	ReeditMessage string
	// --fixup=[(amend|reword):]<commit>
	// Create a new commit which "fixes up" <commit> when applied with git rebase --autosquash.
	Fixup string
	// --squash=<commit>
	// Construct a commit message for use with git rebase --autosquash. The commit message title is taken from the specified commit with a prefix of "squash! ".
	Squash string
	// --reset-author
	// When used with -C/-c/--amend options, declare that the authorship of the resulting commit now belongs to the committer. This also renews the author timestamp.
	ResetAuthor bool
	// --short
	// When doing a dry-run, give the output in the short-format. Implies --dry-run.
	Short bool
	// --branch
	// Show the branch and tracking info even in short-format.
	Branch bool
	// --porcelain
	// When doing a dry-run, give the output in a porcelain-ready format. Implies --dry-run.
	Porcelain bool
	// --long
	// When doing a dry-run, give the output in the long-format. Implies --dry-run.
	Long bool
	// -z
	// --null
	// When showing short or porcelain status output, print the filename verbatim and terminate the entries with NUL, instead of LF.
	Null bool
	// -F <file>
	// --file=<file>
	// Take the commit message from <file>. Use - to read the message from the standard input.
	File string
	// --author=<author>
	// Override the commit author. Specify an explicit author using the standard A U Thor <author@example.com> format.
	Author string
	// --date=<date>
	// Override the author date used in the commit.
	Date string
	// -m <msg>
	// --message=<msg>
	// Use <msg> as the commit message. If multiple -m options are given, their values are concatenated as separate paragraphs.
	Message string
	// -t <file>
	// --template=<file>
	// When editing the commit message, start the editor with the contents in <file>.
	Template string
	// -s
	// --signoff
	// Add a Signed-off-by trailer by the committer at the end of the commit log message.
	Signoff bool
	// --no-signoff
	// Used to countermand an earlier --signoff option.
	NoSignoff bool
	// --trailer <token>[(=|:)<value>]
	// Specify a (<token>, <value>) pair that should be applied as a trailer.
	Trailer string
	// -n
	// --no-verify
	// Bypass the pre-commit and commit-msg hooks. See also githooks[5].
	NoVerify bool
	// --allow-empty
	// Bypasses the safety check that prevents creating a commit with the exact same tree as its parent, primarily for use by scripts.
	AllowEmpty bool
	// --allow-empty-message
	// Create a commit with an empty commit message without using plumbing commands.
	AllowEmptyMessage bool
	// --cleanup=<mode>
	// Determine how the supplied commit message should be cleaned up. The <mode> can be strip, whitespace, verbatim, scissors or default.
	Cleanup string
	// -e
	// --edit
	// Let the user further edit the message taken from another source (e.g., -F, -m, -C).
	Edit bool
	// --no-edit
	// Use the selected commit message without launching an editor.
	NoEdit bool
	// --amend
	// Replace the tip of the current branch by creating a new commit, amending the previous one.
	Amend bool
	// --no-post-rewrite
	// Bypass the post-rewrite hook.
	NoPostRewrite bool
	// -i
	// --include
	// Stage the contents of paths given on the command line before committing.
	Include bool
	// -o
	// --only
	// Make a commit by taking only the updated contents of the paths specified on the command line.
	Only bool
	// --pathspec-from-file=<file>
	// Pass pathspec in <file> instead of commandline args. Use - for standard input.
	PathspecFromFile string
	// --pathspec-file-nul
	// Only with --pathspec-from-file, pathspec elements are separated with NUL character.
	PathspecFileNul bool
	// -u[<mode>]
	// --untracked-files[=<mode>]
	// Show untracked files. The mode can be 'no', 'normal', or 'all'.
	UntrackedFiles string
	// -v
	// --verbose
	// Show unified diff between the HEAD commit and what would be committed in the message template.
	Verbose bool
	// -q
	// --quiet
	// Suppress commit summary message.
	Quiet bool
	// --dry-run
	// Do not create a commit, but show a list of paths that are to be committed.
	DryRun bool
	// --status
	// Include the output of git-status in the commit message template. This is the default.
	Status bool
	// --no-status
	// Do not include the output of git-status in the commit message template.
	NoStatus bool
	// -S[<key-id>]
	// --gpg-sign[=<key-id>]
	// GPG-sign commits. The <key-id> is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space.
	GpgSign string
	// --no-gpg-sign
	// Do not GPG-sign the commit. Useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.
	NoGpgSign bool
	// --
	// Do not interpret any more arguments as options.
	DoNotInterpretMoreArgumentsAsOptions bool
	// <pathspec>...
	// Commit the contents of the files that match the pathspec without recording the changes already added to the index.
	Pathspec []string
}

func CommitCmd(opts *CommitOptions) *exec.Cmd {
	args := []string{"commit"}

	if opts.All {
		args = append(args, "--all")
	}
	if opts.Patch {
		args = append(args, "--patch")
	}
	if opts.ReuseMessage != "" {
		args = append(args, fmt.Sprintf("--reuse-message=%s", opts.ReuseMessage))
	}
	if opts.ReeditMessage != "" {
		args = append(args, fmt.Sprintf("--reedit-message=%s", opts.ReeditMessage))
	}
	if opts.Fixup != "" {
		args = append(args, fmt.Sprintf("--fixup=%s", opts.Fixup))
	}
	if opts.Squash != "" {
		args = append(args, fmt.Sprintf("--squash=%s", opts.Squash))
	}
	if opts.ResetAuthor {
		args = append(args, "--reset-author")
	}
	if opts.Short {
		args = append(args, "--short")
	}
	if opts.Branch {
		args = append(args, "--branch")
	}
	if opts.Porcelain {
		args = append(args, "--porcelain")
	}
	if opts.Long {
		args = append(args, "--long")
	}
	if opts.Null {
		args = append(args, "--null")
	}
	if opts.File != "" {
		args = append(args, fmt.Sprintf("--file=%s", opts.File))
	}
	if opts.Author != "" {
		args = append(args, fmt.Sprintf("--author=%s", opts.Author))
	}
	if opts.Date != "" {
		args = append(args, fmt.Sprintf("--date=%s", opts.Date))
	}
	if opts.Message != "" {
		args = append(args, fmt.Sprintf("--message=%s", opts.Message))
	}
	if opts.Template != "" {
		args = append(args, fmt.Sprintf("--template=%s", opts.Template))
	}
	if opts.Signoff {
		args = append(args, "--signoff")
	}
	if opts.NoSignoff {
		args = append(args, "--no-signoff")
	}
	if opts.Trailer != "" {
		args = append(args, "--trailer")
		args = append(args, opts.Trailer)
	}
	if opts.NoVerify {
		args = append(args, "--no-verify")
	}
	if opts.AllowEmpty {
		args = append(args, "--allow-empty")
	}
	if opts.AllowEmptyMessage {
		args = append(args, "--allow-empty-message")
	}
	if opts.Cleanup != "" {
		args = append(args, fmt.Sprintf("--cleanup=%s", opts.Cleanup))
	}
	if opts.Edit {
		args = append(args, "--edit")
	}
	if opts.NoEdit {
		args = append(args, "--no-edit")
	}
	if opts.Amend {
		args = append(args, "--amend")
	}
	if opts.NoPostRewrite {
		args = append(args, "--no-post-rewrite")
	}
	if opts.Include {
		args = append(args, "--include")
	}
	if opts.Only {
		args = append(args, "--only")
	}
	if opts.PathspecFromFile != "" {
		args = append(args, fmt.Sprintf("--pathspec-from-file=%s", opts.PathspecFromFile))
	}
	if opts.PathspecFileNul {
		args = append(args, "--pathspec-file-nul")
	}
	if opts.UntrackedFiles != "" {
		args = append(args, fmt.Sprintf("--untracked-files=%s", opts.UntrackedFiles))
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.DryRun {
		args = append(args, "--dry-run")
	}
	if opts.Status {
		args = append(args, "--status")
	}
	if opts.NoStatus {
		args = append(args, "--no-status")
	}
	if opts.GpgSign != "" {
		args = append(args, fmt.Sprintf("--gpg-sign=%s", opts.GpgSign))
	}
	if opts.NoGpgSign {
		args = append(args, "--no-gpg-sign")
	}
	if opts.DoNotInterpretMoreArgumentsAsOptions {
		args = append(args, "--")
	}
	if opts.Pathspec != nil {
		args = append(args, opts.Pathspec...)
	}

	return execGit(opts.CmdDir, args)
}

func Commit(opts *CommitOptions) ([]byte, error) {
	return run(CommitCmd(opts))
}
