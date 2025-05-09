package ast

import (
	"github.com/xDeFc0nx/yac/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "lej"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotehrVar",
				},
			},
		},
	}
	if program.String() != "lej myVar = anotherVar" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
