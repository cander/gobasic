package gobasic

import (
	"fmt"
	"strconv"
	"strings"
)

type Statement interface {
	Execute() (int, error)
	LineNo() int
	Text() string
}

// for the moment, statements are not polymorphic. We can add that later befind the Statement interface
type statement struct {
	lineNo int
	text   string
}

const NEXT_LINE = -1

func ParseStatement(line string) (Statement, error) {
	toks := strings.Fields(line)
	lineNo, _ := strconv.Atoi(toks[0]) // ATM we assume all digits already

	result := statement{lineNo, line}

	return result, nil
}

func (s statement) Execute() (int, error) {
	fmt.Printf("executing: %s\n", s.text)
	// switch on type of statement b/c we're not doing polymorphism, yet
	return NEXT_LINE, nil
}

func (s statement) LineNo() int {
	return s.lineNo
}

func (s statement) Text() string {
	return s.text
}
