package gobasic

import (
	"fmt"
	"strconv"
	"strings"
)

type Statement interface {
	Execute() (int, error)
}

// for the moment, statements are not polymorphic. We can add that later befind the Statement interface
type statement struct {
	lineNo int
	text   string
}

func ParseStatement(line string) (Statement, error) {
	toks := strings.Fields(line)
	lineNo, _ := strconv.Atoi(toks[0]) // ATM we assume all digits already

	result := statement{lineNo, line}

	return result, nil
}

func (s statement) Execute() (int, error) {
	fmt.Printf("executing: %s", s.text)
	// switch on type of statement b/c we're not doing polymorphism, yet
	return 0, nil
}
