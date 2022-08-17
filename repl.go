package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		if input != "" {
			parseUserCommand(input, intr)
		}
		fmt.Print("gobasic>> ")
	}

	fmt.Println("\n\nBye!")
	return nil
}

func parseUserCommand(cmdLine string, intr gobasic.Interpreter) {
	toks := strings.Fields(cmdLine)
	cmd := strings.ToUpper(toks[0])

	// could re-do this logic to look for the line number in the whole line before splitting the line
	justDigits, _ := regexp.MatchString(`^\d+$`, cmd)
	if justDigits {
		stmt, _ := gobasic.ParseStatement(cmdLine)
		fmt.Printf("created statement %v\n", stmt)
		intr.UpsertLine(stmt)
	} else {
		switch cmd {
		case "DUMP":
			intr.Dump()
		case "RUN":
			intr.Run()
		default:
			fmt.Printf("Unrecognized command - '%s'\n", cmd)
		}
	}
}
