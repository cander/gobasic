package main

import (
	"fmt"

	gobasic "cander.org/gobasic/pkg"
)

func main() {
	fmt.Println("hello world")

	intr := gobasic.NewInterpreter()

	intr.Dump()
}
