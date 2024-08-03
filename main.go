package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Commit represents a git commit
type Commit struct {
	Hash    string
	Parents []string
	Message string
}

func main() {
	// Get the git log data
	output, err := getGitLog()
	if err != nil {
		fmt.Println("Error getting git log:", err)
		os.Exit(1)
	}

	// Parse the git log data
	commits := parseGitLog(output)

	// Draw the git log graph
	drawGraph(commits)
}

func getGitLog() (string, error) {
	cmd := exec.Command("git", "log", "--pretty=format:%H %P %s")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func parseGitLog(log string) []Commit {
	lines := strings.Split(log, "\n")
	commits := make([]Commit, 0, len(lines))

	for _, line := range lines {
		parts := strings.SplitN(line, " ", 3)
		if len(parts) < 3 {
			continue
		}
		hash := parts[0]
		parents := strings.Split(parts[1], " ")
		message := parts[2]
		commits = append(commits, Commit{Hash: hash, Parents: parents, Message: message})
	}

	return commits
}

func drawGraph(commits []Commit) {
	for _, commit := range commits {
		fmt.Printf("%s\n", commit.Hash[:7])
		fmt.Printf("|\n")
		fmt.Printf("|-- %s\n", commit.Message)
		fmt.Printf("|\n")
	}
}
