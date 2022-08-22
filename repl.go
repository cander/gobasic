package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	gobasic "cander.org/gobasic/pkg"
)

const prompt = "gobasic>> "

func main() {
	intr := gobasic.NewInterpreter()

	readLoop(intr)
}

func readLoop(intr gobasic.Interpreter) error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading commands:", err)
			return err
		}

		fmt.Printf("read input: %s\n", input)
		if input != "" {
			if err := parseUserCommand(input, intr); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}

		fmt.Print(prompt)
	}

	fmt.Println("\n\nBye!")
	return nil
}

func parseUserCommand(cmdLine string, intr gobasic.Interpreter) error {
	toks := strings.Fields(cmdLine)
	cmd := strings.ToUpper(toks[0])

	// could re-do this logic to look for the line number in the whole line before splitting the line
	justDigits, _ := regexp.MatchString(`^\d+$`, cmd)
	if justDigits {
		stmt, err := gobasic.ParseStatement(cmdLine)
		if err != nil {
			return err
		}
		intr.UpsertLine(stmt)
	} else {
		switch cmd {
		case "DUMP":
			intr.Dump()
		case "LIST":
			intr.List()
		case "RUN":
			intr.Run()
		default:
			fmt.Printf("Unrecognized command - '%s'\n", cmd)
		}
	}
	return nil
}
