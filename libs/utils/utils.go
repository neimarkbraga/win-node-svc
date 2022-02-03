package utils

import (
	"os/exec"
	"regexp"
)

func PanicOnError(err error)  {
	if err != nil {
		panic(err)
	}
}

func GetNodeJsExePath() (string, error) {
	cmd := exec.Command("where", "node")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	pattern := regexp.MustCompile("[\n\r]")
	nodejs := pattern.ReplaceAllString(string(output), "")

	return nodejs, nil
}