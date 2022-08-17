package gobasic

import "fmt"

type Interpreter struct {
	prog program
}

func NewInterpreter() Interpreter {
	return Interpreter{newProgram()}
}

func (i Interpreter) UpsertLine(stmt Statement) {
	i.prog.upsertStatement(stmt)
}

func (i Interpreter) Run() {
	pc := i.prog.initialize()

	stmt, err := i.prog.fetchStatement(pc)
	nextLineNo, err := stmt.Execute()

	fmt.Printf("nextLineNo = %d, err = %v\n", nextLineNo, err)

	fmt.Println("only executing one statement")
}

func (i Interpreter) Dump() {
	fmt.Printf("Interpreter state: %v\n", i)
	i.prog.dump()
}
