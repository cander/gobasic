package main

import (
	"bufio"
	"fmt"
	"os"

	gobasic "cander.org/gobasic/pkg"
)

func main() {
	fmt.Println("hello world")

	intr := gobasic.NewInterpreter()

	intr.Dump()
	readLoop()
}

func readLoop() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		fmt.Print("gobasic>> ")
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading commands:", err)
			return err
		}

		fmt.Printf("read input: %s\n", input)
	}

	fmt.Println("\n\nBye!")
	return nil
}
