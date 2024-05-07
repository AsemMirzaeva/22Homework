package git

import (
	"bytes"
	"os/exec"
)

func GetUsername()(string, error) {
	cmd := exec.Command("git", "config", "user.name")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}

func GetUserEmail()(string, error) {
	cmd := exec.Command("git", "config", "user.email")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}