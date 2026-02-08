package main

import (
	"bufio"
	"fmt"
	"os"
)

func execInput(input string) error {
	// remove the newline character
	input = input.TrimSuffix(inputm, "\n")

}

func main() {
	// buffer to take in input
	reader := bufio.NewReader(os.Stdin)

	// shell prompt
	for {
		// read current input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		execInput(input)
	}
}
