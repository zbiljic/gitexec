// Code generated for gitexec. DO NOT EDIT.
//
// git-gc - Cleanup unnecessary files and optimize the local repository
//
// Reference: https://git-scm.com/docs/git-gc
package gitexec

import (
	"fmt"
	"os/exec"
)

type GcOptions struct {
	CmdDir string

	// --aggressive
	// More aggressively optimize the repository at the expense of taking much more time.
	Aggressive bool
	// --auto
	// Check whether any housekeeping is required; if not, exit without performing any work.
	Auto bool
	// --detach
	// Run in the background if the system supports it. This option overrides the gc.autoDetach config.
	Detach bool
	// --no-detach
	// Do not run in the background. This option overrides the gc.autoDetach config.
	NoDetach bool
	// --cruft
	// When expiring unreachable objects, pack them separately into a cruft pack instead of storing them as loose objects.
	Cruft bool
	// --no-cruft
	// Do not pack unreachable objects into a cruft pack.
	NoCruft bool
	// --max-cruft-size=<n>
	// When packing unreachable objects into a cruft pack, limit the size of new cruft packs to be at most <n> bytes.
	MaxCruftSize uint64
	// --expire-to=<dir>
	// When packing unreachable objects into a cruft pack, write a cruft pack containing pruned objects to the directory <dir>.
	ExpireTo string
	// --prune=<date>
	// Prune loose objects older than date (default is 2 weeks ago). --prune=now prunes loose objects regardless of their age.
	Prune string
	// --no-prune
	// Do not prune any loose objects.
	NoPrune bool
	// --quiet
	// Suppress all progress reports.
	Quiet bool
	// --force
	// Force git gc to run even if there may be another git gc instance running on this repository.
	Force bool
	// --keep-largest-pack
	// All packs except the largest non-cruft pack, any packs marked with a .keep file, and any cruft pack(s) are consolidated into a single pack.
	KeepLargestPack bool
}

func GcCmd(opts *GcOptions) *exec.Cmd {
	args := []string{"gc"}

	if opts.Aggressive {
		args = append(args, "--aggressive")
	}
	if opts.Auto {
		args = append(args, "--auto")
	}
	if opts.Detach {
		args = append(args, "--detach")
	}
	if opts.NoDetach {
		args = append(args, "--no-detach")
	}
	if opts.Cruft {
		args = append(args, "--cruft")
	}
	if opts.NoCruft {
		args = append(args, "--no-cruft")
	}
	if opts.MaxCruftSize > 0 {
		args = append(args, fmt.Sprintf("--max-cruft-size=%d", opts.MaxCruftSize))
	}
	if opts.ExpireTo != "" {
		args = append(args, fmt.Sprintf("--expire-to=%s", opts.ExpireTo))
	}
	if opts.Prune != "" {
		args = append(args, fmt.Sprintf("--prune=%s", opts.Prune))
	}
	if opts.NoPrune {
		args = append(args, "--no-prune")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if opts.Force {
		args = append(args, "--force")
	}
	if opts.KeepLargestPack {
		args = append(args, "--keep-largest-pack")
	}

	return execGit(opts.CmdDir, args)
}

func Gc(opts *GcOptions) ([]byte, error) {
	return run(GcCmd(opts))
}
