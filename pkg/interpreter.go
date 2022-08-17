package gobasic

import "fmt"

type Interpreter struct {
	prog program
}

func NewInterpreter() Interpreter {
	return Interpreter{newProgram()}
}

func (i Interpreter) Dump() {
	fmt.Printf("Interpreter state: %v\n", i)
}

func (i Interpreter) UpsertLine(stmt Statement) {
	fmt.Printf("ready to add statement %v\n", stmt)
	i.prog.upsertStatement(stmt)
}
