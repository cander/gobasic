package gobasic

import (
	"errors"
	"fmt"
)

type program struct {
	statements     map[int]Statement
	statementIndex []int
}

type programCounter uint32

const firstPC = 0

var errEndOfProgram = errors.New("end of program")

func newProgram() program {
	return program{make(map[int]Statement), nil}
}

func (p program) upsertStatement(stmt Statement) {
	p.statements[stmt.LineNo()] = stmt
}

func (p *program) initialize() programCounter {
	p.statementIndex = make([]int, 0, len(p.statements))
	// TODO - sort the line numbers before indexing
	for lineNo := range p.statements {
		p.statementIndex = append(p.statementIndex, lineNo)
	}

	return firstPC
}

func (p program) fetchStatement(pc programCounter) (Statement, error) {
	if int(pc) >= len(p.statementIndex) {
		return nil, fmt.Errorf("invalid PC - %d too large", pc)
	}
	lineNo := p.statementIndex[pc]
	result, ok := p.statements[lineNo]
	if !ok {
		return nil, fmt.Errorf("line number %d does not exist at PC %d", lineNo, pc)
	}

	return result, nil
}

func (p program) nextPC(currentPc programCounter, jumpLoc int) (programCounter, error) {
	if jumpLoc == NEXT_LINE {
		currentPc++
		if int(currentPc) >= p.programSize() {
			// end of the program
			return 0, errEndOfProgram
		}
		return currentPc, nil
	} else {
		panic("GOTO not supported - yet")
	}
}

func (p program) programSize() int {
	return len(p.statementIndex)
}

func (p program) dump() {
	fmt.Printf("Program state: %v\n", p)
}
