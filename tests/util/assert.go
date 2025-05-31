package util

import (
	"testing"

	"github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

func AssertTokens(t *testing.T, tok *tokenizer.Tokenizer, expected []tokenizer.Token) {
	actual := tok.Tokens()
	minLen := len(expected)
	if len(actual) < minLen {
		minLen = len(actual)
	}

	for i := 0; i < minLen; i++ {
		if expected[i].Type != actual[i].Type || expected[i].Value != actual[i].Value {
			t.Fatalf("Expected %s, got %s", expected[i].String(), actual[i].String())
		}
	}

	for i := minLen; i < len(expected); i++ {
		t.Logf("Expected token: %v", expected[i])
	}

	for i := minLen; i < len(actual); i++ {
		t.Logf("Unexpected token: %v", actual[i])
	}

	if len(actual) != len(expected) {
		t.Fatalf("Expected %d tokens, got %d", len(expected), len(actual))
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
