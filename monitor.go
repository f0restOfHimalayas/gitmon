package gitmon

import (
	"fmt"
	"os/exec"
	"regexp"

	cfg "github.com/gookit/config/v2"
)

func LoadConfig() ([]string, error) {
	err := cfg.LoadFiles("config.json")
	if err != nil {
		return nil, err
	}
	return cfg.Strings("projectPaths"), nil
}

func FetchLatestCommits(repoPath string) (string, string, error) {
	cmd := exec.Command("git", "fetch")
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return "", "", err
	}

	cmd = exec.Command("git", "log", "origin/master", "-1")
	cmd.Dir = repoPath
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return "", "", err
	}

	fmt.Println(fmt.Sprintf("\nRepo: %v\n%s\n", repoPath, string(output)))

	if len(output) == 0 {
		fmt.Println("no new updates")
		return "", "", err
	}

	commit := string(output)
	return extractCommitID(commit), commit, err
}

func extractCommitID(commitMessage string) string {
	re := regexp.MustCompile(`commit ([0-9a-f]{40})`)
	match := re.FindStringSubmatch(commitMessage)
	if len(match) >= 2 {
		return match[1]
	}
	return ""
}
