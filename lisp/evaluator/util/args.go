package util

import (
	"fmt"

	"github.com/Olian04/go-lisp/lisp/parser/ast"
)

func AssertExactArgs(args []ast.Statement, expected int) {
	if len(args) != expected {
		panic(fmt.Errorf("expected %d arguments, got %d", expected, len(args)))
	}
}

func AssertAtLeastArgs(args []ast.Statement, expected int) {
	if len(args) < expected {
		panic(fmt.Errorf("expected at least %d arguments, got %d", expected, len(args)))
	}
}

func AssertAtMostArgs(args []ast.Statement, expected int) {
	if len(args) > expected {
		panic(fmt.Errorf("expected at most %d arguments, got %d", expected, len(args)))
	}
}
