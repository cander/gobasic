package gobasic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fetchStatement_success(t *testing.T) {
	prog := newProgram()
	stmt100, err := ParseStatement("100 print f")
	assert.NoError(t, err, "parse valid statement") // maybe remove
	prog.upsertStatement(stmt100)

	pc := prog.initialize()

	foundStmt, err := prog.fetchStatement(pc)
	assert.NoError(t, err, "fetch existing statement")
	assert.NotNil(t, foundStmt, "fetch existing statement")
	assert.Equal(t, 100, foundStmt.LineNo(), "incorrect line number")
	// assert.Equal(t, "100 print f", foundStmt.Text(), "incorrect statement text")
}

func Test_fetchStatement_not_found(t *testing.T) {
	prog := newProgram()

	pc := prog.initialize()

	foundStmt, err := prog.fetchStatement(pc)
	assert.ErrorContains(t, err, "invalid PC", "fetch non-existent statement")
	assert.Nil(t, foundStmt, "fetch non-existent statement")
}

func Test_statement_ordering(t *testing.T) {
	prog := newProgram()
	stmt20, err := ParseStatement("20 print twenty")
	assert.NoError(t, err, "parse valid statement") // maybe remove
	prog.upsertStatement(stmt20)
	stmt10, err := ParseStatement("10 print ten")
	assert.NoError(t, err, "parse valid statement") // maybe remove
	prog.upsertStatement(stmt10)
	prog.initialize()

	stmts := prog.listStatements()
	assert.Equal(t, 2, len(stmts), "two statements in list")
	assert.Equal(t, 10, stmts[0].LineNo(), "first statement is 10")
	assert.Equal(t, 20, stmts[1].LineNo(), "second statement is 20")
}

func loadProgram(t *testing.T, lines []string) program {
	result := newProgram()
	for _, line := range lines {
		stmt, err := ParseStatement(line)
		assert.NoError(t, err, "invalid program line")
		result.upsertStatement(stmt)
	}
	result.initialize()
	return result
}

func TestNextPc(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		currentPc programCounter
		jumpLoc   int
		wantErr   bool
		wantPc    programCounter
	}{
		{"simple next", []string{"10 let a = 5", "20 let b = a"}, firstPC, NEXT_LINE, false, 1},
		{"program end", []string{"10 let a = 5", "20 let b = a"}, 1, NEXT_LINE, true, 0},

		{"simple goto", []string{"10 goto 30", "20 let b = a", "30 print b"}, firstPC, 30, false, 2},
		{"invalid goto dest", []string{"10 goto 30"}, firstPC, 30, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prog := loadProgram(t, tt.lines)
			nextPc, err := prog.nextPC(tt.currentPc, tt.jumpLoc)
			if tt.wantErr {
				assert.NotNil(t, err, "expected an err")
			} else {
				assert.Equal(t, tt.wantPc, nextPc, "incorrect PC")
			}
		})
	}
}
