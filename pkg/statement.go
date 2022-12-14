package gobasic

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopl.io/ch7/eval"
)

type Statement interface {
	Execute(env eval.Env) (int, error)
	LineNo() int
	String() string
}

// for the moment, statements are not polymorphic. We can add that later befind the Statement interface
// it would make sense for different types of statements to have different data - currently we assume
// that "rest" can work for all statements.
type statement struct {
	lineNo int
	opCode string
	rest   string
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
	var err error

	switch opCode {
	case "GOTO":
		stmt, err = parseGoto(lineNo, rest)
	case "LET":
		stmt, err = parseLet(lineNo, rest)
	case "INPUT":
		stmt, err = parseInput(lineNo, rest)
	case "PRINT":
		stmt, err = parsePrint(lineNo, rest)
	default:
		return nil, fmt.Errorf("invalid opcode: %s", opCode)
	}

	return stmt, err
}

func (s statement) Execute(env eval.Env) (int, error) {
	fmt.Printf("Error executing: %s\n", s.String())

	return 0, fmt.Errorf("unrecognized statement: %v", s)
}

func (s statement) LineNo() int {
	return s.lineNo
}

func (s statement) String() string {
	return fmt.Sprintf("%5d %s %s", s.lineNo, s.opCode, s.rest)
}

// GOTO

type gotoStatement struct {
	statement
	destLineNo int
}

func parseGoto(lineNo int, rest string) (*gotoStatement, error) {
	var result gotoStatement
	rest = strings.TrimSpace(rest)
	_, err := regexp.MatchString(`^\s*\d+\s*$`, rest)
	if err == nil {
		destLineNo, err := strconv.Atoi(rest)
		if err == nil {
			result = gotoStatement{statement{lineNo, "GOTO", rest}, destLineNo}
		} else {
			return nil, fmt.Errorf("invalid line number: %s", rest)
		}
	}

	return &result, nil
}

func (g gotoStatement) Execute(env eval.Env) (int, error) {
	return g.destLineNo, nil
}

// INPUT

type inputStatement struct {
	statement
	varName eval.Var
}

func parseInput(lineNo int, rest string) (*inputStatement, error) {
	foundVar, err := regexp.MatchString(`^[[:alpha:]][[:alnum:]]*$`, rest)
	if !foundVar || err != nil {
		return nil, fmt.Errorf("invalid variable name: %s", rest)
	}
	varName := eval.Var(rest)
	result := inputStatement{statement{lineNo, "LET", rest}, varName}

	return &result, nil
}

func (i inputStatement) Execute(env eval.Env) (int, error) {
	var inputValue float64 = 0
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputStr := scanner.Text()
		if err := scanner.Err(); err != nil {
			return 0, fmt.Errorf("error reading INPUT: %v", err)
		}

		inputStr = strings.TrimSpace(inputStr)
		if len(inputStr) > 0 {
			inputValue, err = strconv.ParseFloat(inputStr, 64)
			if err != nil {
				return 0, fmt.Errorf("error converting INPUT: %v", err)
			}
			break
		}
	}

	env[i.varName] = inputValue

	return NEXT_LINE, nil
}

// LET

type letStatement struct {
	statement
	varName eval.Var
	letRhs  eval.Expr
}

func parseLet(lineNo int, rest string) (*letStatement, error) {
	// rest: A = 5
	toks := strings.Fields(rest)
	// should check num toks
	if len(toks) < 3 {
		return nil, fmt.Errorf("invalid LET statement: LET %s", rest)
	}
	foundVar, err := regexp.MatchString(`^[[:alpha:]][[:alnum:]]*$`, toks[0])
	if !foundVar || err != nil {
		return nil, fmt.Errorf("invalid variable name: %s", toks[0])
	}
	varName := eval.Var(toks[0])
	if toks[1] != "=" {
		return nil, fmt.Errorf("invalid LET statement: %s", rest)
	}
	lhsExp := regexp.MustCompile(`^.*=\s+`) // everything up to equals
	rhs := lhsExp.ReplaceAllString(rest, "")

	rhsExpr, err := eval.Parse(rhs)
	if err != nil {
		return nil, fmt.Errorf("invalid LET statement: %v", err)
	}

	result := letStatement{statement{lineNo, "LET", rest}, varName, rhsExpr}

	return &result, nil
}

func (l letStatement) Execute(env eval.Env) (int, error) {
	env[l.varName] = l.letRhs.Eval(env)

	return NEXT_LINE, nil
}

// PRINT

type printStatement struct {
	statement
	literalString string
	expr          eval.Expr
}

func parsePrint(lineNo int, rest string) (*printStatement, error) {
	var result printStatement
	re := regexp.MustCompile(`^\s*"([^"]+)"\s*$`)
	quotedStrings := re.FindStringSubmatch(rest)
	if quotedStrings != nil {
		result = printStatement{statement{lineNo, "PRINT", rest}, quotedStrings[1], nil}
	} else {
		expr, err := eval.Parse(rest)
		if err == nil {
			result = printStatement{statement{lineNo, "PRINT", rest}, "", expr}
		} else {
			return nil, err
		}
	}

	return &result, nil
}

func (p printStatement) Execute(env eval.Env) (int, error) {
	if p.expr != nil {
		fmt.Println(p.expr.Eval(env))
	} else {
		fmt.Println(p.literalString)
	}
	return NEXT_LINE, nil
}
