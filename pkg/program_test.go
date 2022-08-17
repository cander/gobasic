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
	assert.Equal(t, "100 print f", foundStmt.Text(), "incorrect statement text")
}

func Test_fetchStatement_not_found(t *testing.T) {
	prog := newProgram()

	pc := prog.initialize()

	foundStmt, err := prog.fetchStatement(pc)
	assert.ErrorContains(t, err, "invalid PC", "fetch non-existent statement")
	assert.Nil(t, foundStmt, "fetch non-existent statement")
}
