package gobasic

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Statement interface {
	Execute() (int, error)
	LineNo() int
	Text() string
	String() string
}

// for the moment, statements are not polymorphic. We can add that later befind the Statement interface
// it would make sense for different types of statements to have different data - currently we assume
// that "rest" can work for all statements.
type statement struct {
	lineNo    int
	opCode    string
	wholeLine string
	rest      string
}

const NEXT_LINE = -1

func ParseStatement(line string) (Statement, error) {
	toks := strings.Fields(line)
	if len(toks) < 3 {
		return nil, fmt.Errorf("syntax error - too few tokens: %s", line)
	}
	lineNo, _ := strconv.Atoi(toks[0]) // ATM we assume all digits already
	opCode := strings.ToUpper(toks[1])
	lineOpExp := regexp.MustCompile(`^.*` + toks[0] + `\s+` + toks[1] + `\s+`)
	rest := lineOpExp.ReplaceAllString(line, "")
	var stmt Statement

	switch opCode {
	case "PRINT":
		stmt = statement{lineNo, opCode, line, rest}
	default:
		return nil, fmt.Errorf("invalid opcode: %s", opCode)
	}

	return stmt, nil
}

func (s statement) Execute() (int, error) {
	fmt.Printf("executing: %s\n", s.wholeLine)
	// switch on type of statement b/c we're not doing polymorphism, yet
	switch s.opCode {
	case "PRINT":
		fmt.Println(s.rest)
	default:
		fmt.Println("unrecognized op code - shouldn't happen")
		//throw an error
	}

	return NEXT_LINE, nil
}

func (s statement) LineNo() int {
	return s.lineNo
}

func (s statement) Text() string {
	return s.wholeLine
}

func (s statement) String() string {
	return fmt.Sprintf("%5d %s %s", s.lineNo, s.opCode, s.rest)
}
