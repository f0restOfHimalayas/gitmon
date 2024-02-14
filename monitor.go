package gitmon

import (
	"fmt"
	"os/exec"

	cfg "github.com/gookit/config/v2"
)

func LoadConfig() ([]string, error) {
	err := cfg.LoadFiles("config.json")
	if err != nil {
		return nil, err
	}
	return cfg.Strings("projectPaths"), nil
}

func FetchLatestCommits(repoPath string) {
	cmd := exec.Command("git", "fetch")
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	cmd = exec.Command("git", "log", "origin/master", "-1")
	cmd.Dir = repoPath
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Print the output of the command
	fmt.Println(fmt.Sprintf("\nRepo: %v,\nUpdate: %s\n", repoPath, string(output)))
}
