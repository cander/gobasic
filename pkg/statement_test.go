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
		{"simple print", "10 print hi", "printStatement", false},
		{"simple let", "10 let a = 69", "letStatement", false},

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

func TestParseLet(t *testing.T) {
	tests := []struct {
		name    string
		rest    string
		wantLhs string
		wantErr bool
	}{
		{"simple let", "a = 69", "a", false},

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
