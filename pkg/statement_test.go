package gobasic

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopl.io/ch7/eval"
)

func TestParse(t *testing.T) {}

func TestParseStatement(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		wantType string
		wantErr  bool
	}{
		{"simple goto", "10 goto 20", "gotoStatement", false},
		{"simple let", "10 let a = 69", "letStatement", false},
		{"simple input", "10 input a", "inputStatement", false},
		{"simple print", "10 print a", "printStatement", false},

		{"no opcode", "10", "", true},
		{"invalid opcode", "10 BARF", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, err := ParseStatement(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				stmtType := reflect.TypeOf(stmt).String()
				assert.Equal(t, "*gobasic."+tt.wantType, stmtType, "incorrect statement type")
			}
		})
	}
}

func TestParseGoto(t *testing.T) {
	tests := []struct {
		name           string
		rest           string
		wantErr        bool
		wantDestLineNo int
	}{
		{"valid goto", "20", false, 20},
		{"trailing space", "20 ", false, 20},

		{"invalid destination", "hell", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, err := parseGoto(100, tt.rest)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, "*gobasic.gotoStatement", reflect.TypeOf(stmt).String(), "not a GOTO statement")

				assert.Equal(t, tt.wantDestLineNo, stmt.destLineNo, "literal string")

			}
		})
	}
}

func TestParseLet(t *testing.T) {
	tests := []struct {
		name    string
		rest    string
		wantLhs string
		wantErr bool
	}{
		{"simple let", "a = 69", "a", false},
		{"many chars let", "cat = 69", "cat", false},
		{"alphanumeric let", "a42 = 69", "a42", false},

		{"backwards", "9 = A", "", true},
		{"incomplete 1", "10", "", true},
		{"incomplete 2", "a =", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, err := parseLet(100, tt.rest)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, "*gobasic.letStatement", reflect.TypeOf(stmt).String(), "not a LET statement")
				assert.Equal(t, eval.Var(tt.wantLhs), stmt.varName, "LHS varaible name")
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		name    string
		rest    string
		wantVar string
		wantErr bool
	}{
		{"simple input", "a", "a", false},

		{"extra garbage", "A = 69", "", true},
		{"invalid var name", "10", "", true},
		{"incomplete", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, err := parseInput(100, tt.rest)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, "*gobasic.inputStatement", reflect.TypeOf(stmt).String(), "not an INPUT statement")
				assert.Equal(t, eval.Var(tt.wantVar), stmt.varName, "varaible name")
			}
		})
	}
}

func TestParsePrint(t *testing.T) {
	tests := []struct {
		name        string
		rest        string
		wantErr     bool
		wantLiteral string
		wantExpr    bool
	}{
		{"print literal", `"hello world"`, false, "hello world", false},
		{"print expr", `A + 5`, false, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt, err := parsePrint(100, tt.rest)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePrint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, "*gobasic.printStatement", reflect.TypeOf(stmt).String(), "not a PRINT statement")
				if tt.wantExpr {
					assert.NotNil(t, stmt.expr, "expression")
				} else {
					assert.Equal(t, tt.wantLiteral, stmt.literalString, "literal string")
				}

			}
		})
	}
}
