package gobasic

import "fmt"

type Interpreter struct{}

func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (i Interpreter) Dump() {
	fmt.Printf("Interpreter state: %v\n", i)
}

func (i Interpreter) UpsertLine(lineNo int, line string) {
	fmt.Printf("ready to add line number %d\n", lineNo)
}
