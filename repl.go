package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	gobasic "cander.org/gobasic/pkg"
)

func main() {
	fmt.Println("hello world")

	intr := gobasic.NewInterpreter()

	readLoop(intr)
}

func readLoop(intr gobasic.Interpreter) error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("gobasic>> ")
	for scanner.Scan() {
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading commands:", err)
			return err
		}

		fmt.Printf("read input: %s\n", input)
		parseUserCommand(input, intr)
		fmt.Print("gobasic>> ")
	}

	fmt.Println("\n\nBye!")
	return nil
}

func parseUserCommand(cmdLine string, intr gobasic.Interpreter) {
	toks := strings.Fields(cmdLine)
	cmd := strings.ToUpper(toks[0])

	switch cmd {
	case "DUMP":
		intr.Dump()
	default:
		fmt.Printf("Unrecognized command - '%s'\n", cmd)
	}
}
