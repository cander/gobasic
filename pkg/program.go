package gobasic

import (
	"errors"
	"fmt"
	"sort"
)

type program struct {
	statements map[int]Statement
	// zero-based array mapping PC locations (0 to N) to the line number at that location
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
	for lineNo := range p.statements {
		p.statementIndex = append(p.statementIndex, lineNo)
	}
	sort.Ints(p.statementIndex)

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
		// this is a gross linear search, but you shouldn't be using GOTOs
		for pc, lineNo := range p.statementIndex {
			if lineNo == jumpLoc {
				return programCounter(pc), nil
			}
		}
		return 0, fmt.Errorf("GOTO - line not found: %d", jumpLoc)
	}
}

func (p program) programSize() int {
	return len(p.statementIndex)
}

func (p program) listStatements() []Statement {
	result := make([]Statement, 0, len(p.statements))

	for _, lineNo := range p.statementIndex {
		result = append(result, p.statements[lineNo])
	}

	return result
}

func (p program) dump() {
	fmt.Printf("Program state: %v\n", p)
}
