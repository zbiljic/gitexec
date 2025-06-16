// Package gitexec provides a Go wrapper around Git command-line operations.
//
// This package offers a type-safe and structured approach to executing Git
// commands from Go applications. It abstracts the complexities of building
// command-line arguments and handling platform-specific execution details,
// while providing comprehensive options structures that mirror Git's
// command-line interface.
//
// Each Git command is represented by a dedicated function with a corresponding
// Options struct that allows for precise configuration of command parameters.
// The package handles command execution, output collection, and error handling
// in a consistent way.
//
// The package handles platform-specific execution details through build
// constraints, ensuring proper process management on both Unix and Windows
// systems.
package gitexec

//go:generate /bin/sh -c "cd ./internal/generator && go run main.go -output ../.. ../../docs/git"
