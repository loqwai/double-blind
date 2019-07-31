package study

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Study is double blind study. Instances of a study are called trials.
type Study struct {
	Name   string  `json:"name"`
	Groups []Group `json:"groups"`
}

// Group represents a study group. A study will generally have a Control Group
// and one or more Experiment Groups.
type Group struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

// Run picks a group at random and executes it. After the command has
// returned, it will prompt the user which group they think they are in.
func (s *Study) Run() (string, error) {
	rand.Seed(time.Now().UnixNano())
	group := s.Groups[rand.Intn(len(s.Groups))]
	command := exec.Command("/usr/bin/env", "bash", "-c", group.Command)
	err := command.Run()
	if err != nil {
		return "", err
	}

	fmt.Println("What group do you think you were in?")
	for i, group := range s.Groups {
		fmt.Printf("%v: %v\n", i, group.Name)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return "", err
	}

	if input < 0 || len(s.Groups) <= input {
		return "", fmt.Errorf("Guess is out of range")
	}

	guess := s.Groups[input]
	if guess == group {
		return "correct", nil
	}
	return "incorrect", nil
}
