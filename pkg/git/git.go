package git

import (
	"os/exec"
	"strings"
)

// Description:
//
//	Attempts to get the current commit hash of the local repository within the current working directory.
//
// Returns:
//
//	The current commit hash if successful, an error otherwise.
func GetCurrentCommitHash() (string, error) {
	output, err := exec.Command("git", "rev-list", "-1", "HEAD").Output()

	if err != nil {
		return "", err
	}

	value := string(output)
	return strings.TrimSpace(value), nil
}
