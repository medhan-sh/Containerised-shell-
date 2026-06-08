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
		target := args[0]
		if target == "~" {
			target = os.Getenv("HOME")
		}
		err := os.Chdir(target)
		if err != nil {
			fmt.Printf("cd: %v: No such file or directory\n", target)
		}
		return os.Chdir(target)
	}

	_, err := exec.LookPath(binary)
	if err != nil {
		fmt.Printf("%v: command not found\n", binary)
	}

	return cmd.Run()

}
func main() {
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
