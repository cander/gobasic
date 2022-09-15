package gobasic

import (
	"fmt"

	"gopl.io/ch7/eval"
)

type Interpreter struct {
	prog program
	env  eval.Env
}

func NewInterpreter() Interpreter {
	return Interpreter{newProgram(), eval.Env{}}
}

func (i Interpreter) UpsertLine(stmt Statement) {
	i.prog.upsertStatement(stmt)
}

func (i Interpreter) Run() { // need to return an error?
	pc := i.prog.initialize()

	for {
		stmt, _ := i.prog.fetchStatement(pc)
		// TODO: handle this error - panic?
		nextLineNo, err := stmt.Execute(i.env)

		fmt.Printf("nextLineNo = %d, err = %v\n", nextLineNo, err)
		pc, err = i.prog.nextPC(pc, nextLineNo)
		if err == errEndOfProgram {
			fmt.Println("Program complete!")
			break
		}
	}
}

func (i Interpreter) List() {
	i.prog.initialize() // this is kinda kludgey - just re-index after each upsert
	fmt.Println("listing...")
	for _, s := range i.prog.listStatements() {
		fmt.Println(s)
	}
}

func (i Interpreter) Dump() {
	fmt.Printf("Interpreter state: %v\n", i)
	i.prog.dump()
	fmt.Println("Varaibles:")
	fmt.Println(i.env)
}
