package gitexec

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestExecGitUsesCmdContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := execGit(ctx, "/tmp/repo", []string{"status", "--short", "--branch"})

	if got, want := cmd.Dir, "/tmp/repo"; got != want {
		t.Fatalf("cmd.Dir = %q, want %q", got, want)
	}

	if cmd.Cancel == nil {
		t.Fatal("expected command created with context to expose Cancel")
	}

	wantArgs := []string{"git", "status", "--short", "--branch"}
	if !reflect.DeepEqual(cmd.Args, wantArgs) {
		t.Fatalf("cmd.Args = %#v, want %#v", cmd.Args, wantArgs)
	}
}

func TestCommandContextCancelsRunningProcess(t *testing.T) {
	fakeGitDir := t.TempDir()
	markerPath := filepath.Join(fakeGitDir, "started")
	writeFakeGit(t, fakeGitDir, markerPath)

	t.Setenv("PATH", fakeGitDir)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := CommandContext(ctx, t.TempDir(), "status")
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected command to be canceled")
	}

	if ctx.Err() == nil {
		t.Fatal("expected context deadline to be exceeded")
	}

	if _, statErr := os.Stat(markerPath); statErr != nil {
		t.Fatalf("expected fake git to start: %v (command error: %v)", statErr, err)
	}

	if elapsed >= 2*time.Second {
		t.Fatalf("expected cancellation before sleep completed, took %s", elapsed)
	}
}

func writeFakeGit(t *testing.T, dir, markerPath string) {
	t.Helper()

	name := "git"
	content := "#!/bin/sh\n" +
		"printf started > " + quotePath(markerPath) + "\n" +
		"/bin/sleep 5\n"

	if runtime.GOOS == "windows" {
		name = "git.bat"
		content = "@echo off\r\n" +
			"echo started > " + quotePath(markerPath) + "\r\n" +
			"timeout /T 5 /NOBREAK >NUL\r\n"
	}

	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o755); err != nil {
		t.Fatalf("write fake git: %v", err)
	}
}

func quotePath(path string) string {
	return `"` + strings.ReplaceAll(path, `"`, `\"`) + `"`
}
