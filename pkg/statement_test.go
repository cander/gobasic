package gobasic

import (
	"reflect"
	"testing"
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
				if stmtType != ("gobasic." + tt.wantType) {
					t.Errorf("Expected Statement type %s, found %s", tt.wantType, stmtType)
				}
			}
		})
	}
}
