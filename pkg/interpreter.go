package gobasic

import "fmt"

type Interpreter struct{}

func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (i Interpreter) Dump() {
	fmt.Printf("Interpreter state: %v\n", i)
}
