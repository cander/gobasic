package gobasic

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gopl.io/ch7/eval"
)

type Statement interface {
	Execute(env eval.Env) (int, error)
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
	letVar    eval.Var
	letRhs    eval.Expr
}

const NEXT_LINE = -1

// yeah, this "parsing" is very hacky - i.e., there is no lexer - it's just a bunch of string
// hackery with regexps, etc. Ultimately, it's left over from the original Bourne Basic implementation,
// or at least that's my story for the moment.
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
	case "LET":
		// 10 LET A = 5
		// should check num toks
		foundVar, err := regexp.MatchString(`^[[:alpha:]][[:alnum:]]*$`, toks[2])
		if !foundVar || err != nil {
			return nil, fmt.Errorf("invalid variable name: %s", toks[2])
		}
		varName := eval.Var(toks[2])
		if toks[3] != "=" {
			return nil, fmt.Errorf("invalid LET statement: %s", rest)
		}
		lhsExp := regexp.MustCompile(`^.*=\s+`) // everything up to equals
		rhs := lhsExp.ReplaceAllString(line, "")

		rhsExpr, err := eval.Parse(rhs)
		if err != nil {
			return nil, fmt.Errorf("invalid LET statement: %v", err)
		}
		fmt.Printf("parsed LET %s = %v\n", varName, rhsExpr)
		stmt = statement{lineNo, opCode, line, rest, varName, rhsExpr}
	case "PRINT":
		stmt = statement{lineNo, opCode, line, rest, "", nil}
	default:
		return nil, fmt.Errorf("invalid opcode: %s", opCode)
	}

	return stmt, nil
}

func (s statement) Execute(env eval.Env) (int, error) {
	fmt.Printf("executing: %s\n", s.wholeLine)
	// switch on type of statement b/c we're not doing polymorphism, yet
	switch s.opCode {
	case "LET":
		env[s.letVar] = s.letRhs.Eval(env)
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
