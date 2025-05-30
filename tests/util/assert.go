package util

import (
	"testing"

	"github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

func AssertNextToken(t *testing.T, tok *tokenizer.Tokenizer, token tokenizer.Token) {
	actual := tok.NextToken()
	if actual.Type != token.Type || actual.Value != token.Value {
		t.Fatalf("Expected %s, got %s", token.String(), actual.String())
	}
}

func AssertProgram(t *testing.T, program ast.Program, expected []ast.Statement) {
	if len(program.Statements) != len(expected) {
		t.Fatalf("Expected %d statements, got %d", len(expected), len(program.Statements))
	}
	for i, statement := range program.Statements {
		if statement.String() != expected[i].String() {
			t.Fatalf("Expected %s, got %s", expected[i].String(), statement.String())
		}
	}
}
