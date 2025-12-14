package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	err := cmd.Run()
	return err == nil
}

func HasStagedChanges() (bool, error) {
	cmd := exec.Command("git", "diff", "--cached", "--quiet")
	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// Exit code 1 means there are changes
			return exitErr.ExitCode() == 1, nil
		}
		return false, fmt.Errorf("failed to check staged changes: %w", err)
	}
	// Exit code 0 means no changes
	return false, nil
}

func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git diff failed: %s", stderr.String())
	}

	return stdout.String(), nil
}

func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("git diff --name-only failed: %s", stderr.String())
	}

	output := strings.TrimSpace(stdout.String())
	if output == "" {
		return []string{}, nil
	}

	return strings.Split(output, "\n"), nil
}

func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git commit failed: %s", stderr.String())
	}

	return nil
}
