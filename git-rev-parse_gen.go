// Code generated for gitexec. DO NOT EDIT.
//
// git-rev-parse - Pick out and massage parameters
//
// Reference: https://git-scm.com/docs/git-rev-parse
package gitexec

import (
	"fmt"
	"os/exec"
)

type RevParseOptions struct {
	CmdDir string

	// --parseopt
	// Use git rev-parse in option parsing mode (see PARSEOPT section below).
	Parseopt bool
	// --sq-quote
	// Use git rev-parse in shell quoting mode (see SQ-QUOTE section below). In contrast to the --sq option below, this mode only does quoting. Nothing else is done to command input.
	SqQuote bool
	// --keep-dashdash
	// Only meaningful in --parseopt mode. Tells the option parser to echo out the first -- met instead of skipping it.
	KeepDashdash bool
	// --stop-at-non-option
	// Only meaningful in --parseopt mode. Lets the option parser stop at the first non-option argument. This can be used to parse sub-commands that take options themselves.
	StopAtNonOption bool
	// --stuck-long
	// Only meaningful in --parseopt mode. Output the options in their long form if available, and with their arguments stuck.
	StuckLong bool
	// --revs-only
	// Do not output flags and parameters not meant for git rev-list command.
	RevsOnly bool
	// --no-revs
	// Do not output flags and parameters meant for git rev-list command.
	NoRevs bool
	// --flags
	// Do not output non-flag parameters.
	Flags bool
	// --no-flags
	// Do not output flag parameters.
	NoFlags bool
	// --default <arg>
	// If there is no parameter given by the user, use <arg> instead.
	Default string
	// --prefix <arg>
	// Behave as if git rev-parse was invoked from the <arg> subdirectory of the working tree. Any relative filenames are resolved as if they are prefixed by <arg> and will be printed in that form.
	Prefix string
	// --verify
	// Verify that exactly one parameter is provided, and that it can be turned into a raw 20-byte SHA-1 that can be used to access the object database. If so, emit it to the standard output; otherwise, error out.
	Verify bool
	// -q
	// --quiet
	// Only meaningful in --verify mode. Do not output an error message if the first argument is not a valid object name; instead exit with non-zero status silently. SHA-1s for valid object names are printed to stdout on success.
	Quiet bool
	// --sq
	// Usually the output is made one line per flag and parameter. This option makes output a single line, properly quoted for consumption by shell. Useful when you expect your parameter to contain whitespaces and newlines (e.g. when using pickaxe -S with git diff-*). In contrast to the --sq-quote option, the command input is still interpreted as usual.
	Sq bool
	// --short[=<length>]
	// Same as --verify but shortens the object name to a unique prefix with at least length characters. The minimum length is 4, the default is the effective value of the core.abbrev configuration variable (see git-config[1]).
	Short string
	// --not
	// When showing object names, prefix them with ^ and strip ^ prefix from the object names that already have one.
	Not bool
	// --abbrev-ref[=(strict|loose)]
	// A non-ambiguous short name of the objects name. The option core.warnAmbiguousRefs is used to select the strict abbreviation mode.
	AbbrevRef string
	// --symbolic
	// Usually the object names are output in SHA-1 form (with possible ^ prefix); this option makes them output in a form as close to the original input as possible.
	Symbolic bool
	// --symbolic-full-name
	// This is similar to --symbolic, but it omits input that is not a ref name (i.e. branch or tag name).
	SymbolicFullName bool
	// --output-object-format=(sha1|sha256|storage)
	// Allow oids to be input from any object format that the current repository supports.
	OutputObjectFormat string
	// --all
	// Show all refs found in refs/.
	All bool
	// --branches[=<pattern>]
	// Show all branches (i.e., refs found in refs/heads). If a pattern is given, only refs matching the given shell glob are shown. If the pattern does not contain a globbing character ( ?, *, or [), it is turned into a prefix match by appending /*.
	Branches string
	// --tags[=<pattern>]
	// Show all tags (i.e., refs found in refs/tags). If a pattern is given, only refs matching the given shell glob are shown. If the pattern does not contain a globbing character ( ?, *, or [), it is turned into a prefix match by appending /*.
	Tags string
	// --remotes[=<pattern>]
	// Show all remote-tracking branches (i.e., refs found in refs/remotes). If a pattern is given, only refs matching the given shell glob are shown. If the pattern does not contain a globbing character ( ?, *, or [), it is turned into a prefix match by appending /*.
	Remotes string
	// --glob=<pattern>
	// Show all refs matching the shell glob pattern pattern. If the pattern does not start with refs/, this is automatically prepended. If the pattern does not contain a globbing character ( ?, *, or [), it is turned into a prefix match by appending /*.
	Glob string
	// --exclude=<glob-pattern>
	// Do not include refs matching <glob-pattern> that the next --all, --branches, --tags, --remotes, or --glob would otherwise consider. Repetitions of this option accumulate exclusion patterns up to the next --all, --branches, --tags, --remotes, or --glob option (other options or arguments do not clear accumulated patterns). The patterns given should not begin with refs/heads, refs/tags, or refs/remotes when applied to --branches, --tags, or --remotes, respectively, and they must begin with refs/ when applied to --glob or --all. If a trailing /* is intended, it must be given explicitly.
	Exclude string
	// --exclude-hidden=(fetch|receive|uploadpack)
	// Do not include refs that would be hidden by git-fetch, git-receive-pack or git-upload-pack by consulting the appropriate fetch.hideRefs, receive.hideRefs or uploadpack.hideRefs configuration along with transfer.hideRefs (see git-config[1]). This option affects the next pseudo-ref option --all or --glob and is cleared after processing them.
	ExcludeHidden string
	// --disambiguate=<prefix>
	// Show every object whose name begins with the given prefix. The <prefix> must be at least 4 hexadecimal digits long to avoid listing each and every object in the repository by mistake.
	Disambiguate string
	// --local-env-vars
	// List the GIT_* environment variables that are local to the repository (e.g. GIT_DIR or GIT_WORK_TREE, but not GIT_EDITOR). Only the names of the variables are listed, not their value, even if they are set.
	LocalEnvVars bool
	// --path-format=(absolute|relative)
	// Controls the behavior of certain other options. If specified as absolute, the paths printed by those options will be absolute and canonical. If specified as relative, the paths will be relative to the current working directory if that is possible. The default is option specific. This option may be specified multiple times and affects only the arguments that follow it on the command line, either to the end of the command line or the next instance of this option.
	PathFormat string
	// --git-dir
	// Show $GIT_DIR if defined. Otherwise show the path to the .git directory. The path shown, when relative, is relative to the current working directory. If $GIT_DIR is not defined and the current directory is not detected to lie in a Git repository or work tree print a message to stderr and exit with nonzero status.
	GitDir bool
	// --git-common-dir
	// Show $GIT_COMMON_DIR if defined, else $GIT_DIR.
	GitCommonDir bool
	// --resolve-git-dir <path>
	// Check if <path> is a valid repository or a gitfile that points at a valid repository, and print the location of the repository. If <path> is a gitfile then the resolved path to the real repository is printed.
	ResolveGitDir string
	// --git-path <path>
	// Resolve "$GIT_DIR/<path>" and takes other path relocation variables such as $GIT_OBJECT_DIRECTORY, $GIT_INDEX_FILE... into account. For example, if $GIT_OBJECT_DIRECTORY is set to /foo/bar then "git rev-parse --git-path objects/abc" returns /foo/bar/abc.
	GitPath string
	// --show-toplevel
	// Show the (by default, absolute) path of the top-level directory of the working tree. If there is no working tree, report an error.
	ShowToplevel bool
	// --show-superproject-working-tree
	// Show the absolute path of the root of the superproject’s working tree (if exists) that uses the current repository as its submodule. Outputs nothing if the current repository is not used as a submodule by any project.
	ShowSuperprojectWorkingTree bool
	// --shared-index-path
	// Show the path to the shared index file in split index mode, or empty if not in split-index mode.
	SharedIndexPath bool
	// --absolute-git-dir
	// Like --git-dir, but its output is always the canonicalized absolute path.
	AbsoluteGitDir bool
	// --is-inside-git-dir
	// When the current working directory is below the repository directory print "true", otherwise "false".
	IsInsideGitDir bool
	// --is-inside-work-tree
	// When the current working directory is inside the work tree of the repository print "true", otherwise "false".
	IsInsideWorkTree bool
	// --is-bare-repository
	// When the repository is bare print "true", otherwise "false".
	IsBareRepository bool
	// --is-shallow-repository
	// When the repository is shallow print "true", otherwise "false".
	IsShallowRepository bool
	// --show-cdup
	// When the command is invoked from a subdirectory, show the path of the top-level directory relative to the current directory (typically a sequence of "../", or an empty string).
	ShowCdup bool
	// --show-prefix
	// When the command is invoked from a subdirectory, show the path of the current directory relative to the top-level directory.
	ShowPrefix bool
	// --show-object-format[=(storage|input|output)]
	// Show the object format (hash algorithm) used for the repository for storage inside the .git directory, input, or output. For input, multiple algorithms may be printed, space-separated. If not specified, the default is "storage".
	ShowObjectFormat string
	// --show-ref-format
	// Show the reference storage format used for the repository.
	ShowRefFormat bool
	// --since=<datestring>
	// --after=<datestring>
	// Parse the date string, and output the corresponding --max-age= parameter for git rev-list.
	Since string
	// --until=<datestring>
	// --before=<datestring>
	// Parse the date string, and output the corresponding --min-age= parameter for git rev-list.
	Until string
	// <arg>...
	// Parameters to be picked out and massaged. These are typically a mixture of flags and parameters for underlying git commands.
	Arg []string
}

func RevParseCmd(opts *RevParseOptions) *exec.Cmd {
	args := []string{"rev-parse"}

	if opts.Parseopt {
		args = append(args, "--parseopt")
	}
	if opts.SqQuote {
		args = append(args, "--sq-quote")
	}
	if opts.KeepDashdash {
		args = append(args, "--keep-dashdash")
	}
	if opts.StopAtNonOption {
		args = append(args, "--stop-at-non-option")
	}
	if opts.StuckLong {
		args = append(args, "--stuck-long")
	}
	if opts.RevsOnly {
		args = append(args, "--revs-only")
	}
	if opts.NoRevs {
		args = append(args, "--no-revs")
	}
	if opts.Flags {
		args = append(args, "--flags")
	}
	if opts.NoFlags {
		args = append(args, "--no-flags")
	}
	if opts.Default != "" {
		args = append(args, "--default")
		args = append(args, opts.Default)
	}
	if opts.Prefix != "" {
		args = append(args, "--prefix")
		args = append(args, opts.Prefix)
	}
	if opts.Verify {
		args = append(args, "--verify")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Sq {
		args = append(args, "--sq")
	}
	if opts.Short != "" {
		args = append(args, fmt.Sprintf("--short=%s", opts.Short))
	}
	if opts.Not {
		args = append(args, "--not")
	}
	if opts.AbbrevRef != "" {
		args = append(args, fmt.Sprintf("--abbrev-ref=%s", opts.AbbrevRef))
	}
	if opts.Symbolic {
		args = append(args, "--symbolic")
	}
	if opts.SymbolicFullName {
		args = append(args, "--symbolic-full-name")
	}
	if opts.OutputObjectFormat != "" {
		args = append(args, fmt.Sprintf("--output-object-format=%s", opts.OutputObjectFormat))
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
	if opts.Disambiguate != "" {
		args = append(args, fmt.Sprintf("--disambiguate=%s", opts.Disambiguate))
	}
	if opts.LocalEnvVars {
		args = append(args, "--local-env-vars")
	}
	if opts.PathFormat != "" {
		args = append(args, fmt.Sprintf("--path-format=%s", opts.PathFormat))
	}
	if opts.GitDir {
		args = append(args, "--git-dir")
	}
	if opts.GitCommonDir {
		args = append(args, "--git-common-dir")
	}
	if opts.ResolveGitDir != "" {
		args = append(args, "--resolve-git-dir")
		args = append(args, opts.ResolveGitDir)
	}
	if opts.GitPath != "" {
		args = append(args, "--git-path")
		args = append(args, opts.GitPath)
	}
	if opts.ShowToplevel {
		args = append(args, "--show-toplevel")
	}
	if opts.ShowSuperprojectWorkingTree {
		args = append(args, "--show-superproject-working-tree")
	}
	if opts.SharedIndexPath {
		args = append(args, "--shared-index-path")
	}
	if opts.AbsoluteGitDir {
		args = append(args, "--absolute-git-dir")
	}
	if opts.IsInsideGitDir {
		args = append(args, "--is-inside-git-dir")
	}
	if opts.IsInsideWorkTree {
		args = append(args, "--is-inside-work-tree")
	}
	if opts.IsBareRepository {
		args = append(args, "--is-bare-repository")
	}
	if opts.IsShallowRepository {
		args = append(args, "--is-shallow-repository")
	}
	if opts.ShowCdup {
		args = append(args, "--show-cdup")
	}
	if opts.ShowPrefix {
		args = append(args, "--show-prefix")
	}
	if opts.ShowObjectFormat != "" {
		args = append(args, fmt.Sprintf("--show-object-format=%s", opts.ShowObjectFormat))
	}
	if opts.ShowRefFormat {
		args = append(args, "--show-ref-format")
	}
	if opts.Since != "" {
		args = append(args, fmt.Sprintf("--since=%s", opts.Since))
	}
	if opts.Until != "" {
		args = append(args, fmt.Sprintf("--until=%s", opts.Until))
	}
	if opts.Arg != nil {
		args = append(args, opts.Arg...)
	}

	return execGit(opts.CmdDir, args)
}

func RevParse(opts *RevParseOptions) ([]byte, error) {
	return run(RevParseCmd(opts))
}
