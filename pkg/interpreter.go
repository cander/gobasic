package gobasic

import "fmt"

type interpreter struct{}

func NewInterpreter() interpreter {
	return interpreter{}
}

func (i interpreter) Dump() {
	fmt.Printf("interpreter: %v\n", i)
}
