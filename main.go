package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	// remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// split input to separate the command and the arguments
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// 'cd' tp home dir with empty path not yet supported
		if len(args) < 2 {
			return errors.New("path required")
		}
		// change the directory and return the error
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// prepare command to execute
	cmd := exec.Command(args[0], args[1:]...)

	// set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// execute the command and return the error
	return cmd.Run()
}

func main() {
	// buffer to take in input
	reader := bufio.NewReader(os.Stdin)

	// shell prompt
	for {
		fmt.Print("> ")
		// read current input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// handle execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
