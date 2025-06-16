# gitexec

[![Go Reference](https://pkg.go.dev/badge/github.com/zbiljic/gitexec.svg)](https://pkg.go.dev/github.com/zbiljic/gitexec)
[![Go Report Card](https://goreportcard.com/badge/github.com/zbiljic/gitexec)](https://goreportcard.com/report/github.com/zbiljic/gitexec)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`gitexec` is a Go library that provides a type-safe and structured wrapper around the Git command-line interface. It simplifies executing Git commands from Go applications by abstracting away the complexities of command-line argument construction and execution.

## Why `gitexec`?

Interacting with Git from a Go application often means manually constructing arguments for `os/exec`. This can be tedious, error-prone, and hard to maintain, especially for commands with many flags.

`gitexec` solves this by:

- **Providing a Type-Safe API**: Each Git command has a dedicated function and a corresponding `Options` struct. This makes discovering and using Git options as simple as filling in a struct field, with the benefits of compile-time checks and IDE autocompletion.
- **Leveraging the Power of Native Git**: By wrapping the `git` CLI, `gitexec` ensures 100% compatibility with all features, configurations, and authentication methods of your installed Git version. You don't have to worry about a reimplementation's limitations.
- **Simplifying Output Handling**: Each command function returns the combined output (`[]byte`) and an `error`, providing a consistent and straightforward way to handle command results.

## Features

- **Auto-generated Wrappers**: Command wrappers are generated directly from Git's documentation, ensuring accuracy and making it easy to add support for new commands.
- **Comprehensive Option Structs**: Each command's `Options` struct mirrors the available command-line flags, complete with documentation.
- **Simple and Consistent API**: All generated commands follow the same `gitexec.Command(opts)` pattern.
- **Cross-Platform**: Handles process management details for Unix and Windows systems.
- **Generic Command Execution**: A `gitexec.Command()` function is available for running arbitrary or not-yet-generated Git commands.

## Installation

```sh
go get github.com/zbiljic/gitexec
```

## Usage

All you need is a standard Git installation available in your system's `PATH`.

### Example: Get the current status

This example runs `git status --short --branch` in the specified directory.

```go
package main

import (
	"fmt"
	"log"

	"github.com/zbiljic/gitexec"
)

func main() {
	output, err := gitexec.Status(&gitexec.StatusOptions{
		CmdDir: "/path/to/your/repo",
		Short:  true,
		Branch: true,
	})
	if err != nil {
		log.Fatalf("git status failed: %v\nOutput: %s", err, string(output))
	}

	fmt.Println("Git Status:")
	fmt.Println(string(output))
}
```

### Example: Get the last 5 commit logs

This example runs `git log --max-count=5 --oneline --graph`.

```go
package main

import (
	"fmt"
	"log"

	"github.com/zbiljic/gitexec"
)

func main() {
	output, err := gitexec.Log(&gitexec.LogOptions{
		CmdDir:   "/path/to/your/repo",
		MaxCount: 5,
		Oneline:  true,
		Graph:    true,
	})
	if err != nil {
		log.Fatalf("git log failed: %v\nOutput: %s", err, string(output))
	}

	fmt.Println("Recent Commits:")
	fmt.Println(string(output))
}
```

### Example: Running an arbitrary command

For commands that are not yet generated, or to pass arguments in a more dynamic way, you can use the generic `gitexec.Command` function. This example runs `git config core.filemode false` inside a repository.

```go
package main

import (
	"fmt"
	"log"

	"github.com/zbiljic/gitexec"
)

func main() {
	repoPath := "/path/to/your/repo"

	// Use gitexec.Command for any git subcommand.
	// The first argument to Command is the working directory.
	// The second is the command name, followed by its arguments.
	output, err := gitexec.Command(repoPath, "config", "core.filemode", "false")
	if err != nil {
		// 'git config' produces no output on success, so output here is for debugging.
		log.Fatalf("git config failed: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("Successfully set 'core.filemode = false' in %s\n", repoPath)
}
```

## Supported Commands

`gitexec` provides generated, type-safe wrappers for the following Git commands:

- `add`
- `branch`
- `clone`
- `commit`
- `diff`
- `fetch`
- `gc`
- `log`
- `ls-remote`
- `pull`
- `reset`
- `rev-list`
- `rev-parse`
- `status`
- `symbolic-ref`

## Contributing

Contributions are welcome! If you find a bug, have a feature request, or want to improve the codebase, please feel free to open an issue or submit a pull request.

### Adding a New Command

The core of this library is its code generator, which parses JSON definitions of Git commands. To add support for a new command:

1.  Create a JSON file in the `docs/git/` directory that describes the command's name, description, and options. Follow the format of the existing JSON files.
2.  Run the code generator:
    ```sh
    go generate ./...
    # or
    make generate
    ```
3.  This will create a new `git-<command>_gen.go` file in the root directory.
4.  Commit the new JSON file and the generated Go file (following pattern of previous commit messages), and open a pull request.

### Development

The project uses `mise` to manage development tools. To install them, run:

```sh
make bootstrap
```

Common development tasks are available in the `Makefile`:

- `make tidy`: Tidy Go modules.
- `make gofmt`: Format Go code with `gofumpt`.
- `make lint`: Lint the source code.
- `make pre-commit`: Run formatters and linters.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
