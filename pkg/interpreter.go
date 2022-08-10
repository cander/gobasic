package gobasic

import "fmt"

type Interpreter struct{}

func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (i Interpreter) Dump() {
	fmt.Printf("Interpreter state: %v\n", i)
}

func (i Interpreter) UpsertLine(stmt Statement) {
	fmt.Printf("ready to add statement %v\n", stmt)
}
