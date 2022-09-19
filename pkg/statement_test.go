package gobasic

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {}

func TestParseStatement(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    Statement
		wantErr bool
	}{
		// TODO: Add test cases.
		// {"simple print", "100 print hi", printStatement{statement{100, "PRINT", "hi", "", nil}, "HI"}, false},

		{"no opcode", "100", nil, true},
		{"invalid opcode", "100 BARF", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStatement(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStatement() = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}
