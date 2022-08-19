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

func (i Interpreter) Run() { // need to return an error?
	pc := i.prog.initialize()

	for {
		stmt, _ := i.prog.fetchStatement(pc)
		// TODO: handle this error - panic?
		nextLineNo, err := stmt.Execute()

		fmt.Printf("nextLineNo = %d, err = %v\n", nextLineNo, err)
		pc, err = i.prog.nextPC(pc, nextLineNo)
		if err == errEndOfProgram {
			fmt.Println("Program complete!")
			break
		}
	}
}

func (i Interpreter) Dump() {
	fmt.Printf("Interpreter state: %v\n", i)
	i.prog.dump()
}
