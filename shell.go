package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	binary := strings.Fields(input)[0]
	args := strings.Fields(input)[1:]

	cmd := exec.Command(binary, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	switch binary {
	case "exit":
		os.Exit(0)
	case "cd":
		var target string
		if len(args) == 0 {
			target = os.Getenv("HOME")
		}
		if target == "~" {
			target = os.Getenv("HOME")
		}
		if err := os.Chdir(target); err != nil {
			fmt.Printf("cd: %v: No such file or directory\n", target)
		}
		return nil
	}

	_, err := exec.LookPath(binary)
	if err != nil {
		fmt.Printf("%v: command not found\n", binary)
	}

	return cmd.Run()

}
func runShell() {
	reader := bufio.NewReader((os.Stdin))
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error while reading: ", err)
			os.Exit(1)
		}
		if input == "" {
			continue
		}

		if err := execInput(input); err != nil {
			fmt.Fprint(os.Stderr)
		}
	}

}
