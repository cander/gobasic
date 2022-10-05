package gobasic

import (
	"bufio"
	"fmt"
	"os"

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
		if err != nil {
			fmt.Printf("ERROR: failed to execute \"%s\" - %s \n", stmt, err)
			break
		}

		pc, err = i.prog.nextPC(pc, nextLineNo)
		if err == errEndOfProgram {
			fmt.Println("Program complete!")
			break
		} else if err != nil {
			fmt.Printf("ERROR: %s\n", err)
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

func (i *Interpreter) Load(fileName string) error {
	fileIn, err := os.Open(fileName)
	if err != nil {
		return err
	}

	i.Reset()
	scanner := bufio.NewScanner(fileIn)
	for scanner.Scan() {
		line := scanner.Text()
		stmt, err := ParseStatement(line)
		if err == nil {
			i.UpsertLine(stmt)
		} else {
			return err
		}
	}

	return nil
}

func (i *Interpreter) Reset() {
	i.prog = newProgram()
	i.env = eval.Env{}
}
