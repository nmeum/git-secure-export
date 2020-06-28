package export

import (
	"os"
	"os/exec"
)

func min(a int64, b int64) int64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func GetDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Stderr = os.Stderr

	dir, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Strip terminating newline
	dir = dir[:len(dir)-1]

	return string(dir), nil
}
