package utils

import "os/exec"

func Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	return cmd.Run()
}
