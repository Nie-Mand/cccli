package utils

import "os/exec"

func Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	return cmd.Run()
}

func Add(files ...string) error {
	if len(files) == 0 {
		return nil
	}

	if files[0] == "." {
		cmd := exec.Command("git", "add", ".")
		return cmd.Run()
	}

	args := append([]string{"add"}, files...)
	cmd := exec.Command("git", args...)
	return cmd.Run()
}
