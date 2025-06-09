package util

import (
	"testing"

	"github.com/Olian04/go-lisp/lisp/parser/ast"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

type assertBuilder struct {
	tb    testing.TB
	value any
}

func Assert(tb testing.TB, value any) assertBuilder {
	tb.Helper()
	return assertBuilder{tb: tb, value: value}
}

func (a assertBuilder) Equal(expectedValue any) assertBuilder {
	a.tb.Helper()
	if a.value != expectedValue {
		a.tb.Fatalf("expected (%T) %v, but got (%T) %v", expectedValue, expectedValue, a.value, a.value)
	}
	return a
}

func (a assertBuilder) NotEqual(expectedValue any) assertBuilder {
	a.tb.Helper()
	if a.value == expectedValue {
		a.tb.Fatalf("expected not (%T) %v, but got (%T) %v", expectedValue, expectedValue, a.value, a.value)
	}
	return a
}

func (a assertBuilder) True() assertBuilder {
	a.tb.Helper()
	if a.value != true {
		a.tb.Fatalf("expected true, but got (%T) %v", a.value, a.value)
	}
	return a
}

func (a assertBuilder) False() assertBuilder {
	a.tb.Helper()
	if a.value != false {
		a.tb.Fatalf("expected false, but got (%T) %v", a.value, a.value)
	}
	return a
}

func (a assertBuilder) Nil() assertBuilder {
	a.tb.Helper()
	if a.value != nil {
		a.tb.Fatalf("expected (%T) %v to be nil", a.value, a.value)
	}
	return a
}

func (a assertBuilder) NotNil() assertBuilder {
	a.tb.Helper()
	if a.value == nil {
		a.tb.Fatalf("expected (%T) %v to not be nil", a.value, a.value)
	}
	return a
}

func (a assertBuilder) Empty() assertBuilder {
	a.tb.Helper()
	switch a.value.(type) {
	case string:
		if a.value != "" {
			a.tb.Fatalf("expected (%T) %v to be empty", a.value, a.value)
		}
	case []any:
		if len(a.value.([]any)) != 0 {
			a.tb.Fatalf("expected (%T) %v to be empty", a.value, a.value)
		}
	case map[any]any:
		if len(a.value.(map[any]any)) != 0 {
			a.tb.Fatalf("expected (%T) %v to be empty", a.value, a.value)
		}
	}
	return a
}

func (a assertBuilder) NotEmpty() assertBuilder {
	a.tb.Helper()
	switch a.value.(type) {
	case string:
		if a.value == "" {
			a.tb.Fatalf("expected (%T) %v to not be empty", a.value, a.value)
		}
	case []any:
		if len(a.value.([]any)) == 0 {
			a.tb.Fatalf("expected (%T) %v to not be empty", a.value, a.value)
		}
	case map[any]any:
		if len(a.value.(map[any]any)) == 0 {
			a.tb.Fatalf("expected (%T) %v to not be empty", a.value, a.value)
		}
	}
	return a
}

func (a assertBuilder) Error(expectedError error) assertBuilder {
	a.tb.Helper()
	if a.value == nil {
		a.tb.Fatalf("expected error, but got nil")
	}
	if a.value.(error).Error() != expectedError.Error() {
		a.tb.Fatalf("expected error %v, but got %v", expectedError, a.value)
	}
	return a
}

func (a assertBuilder) NotError() assertBuilder {
	a.tb.Helper()
	if a.value != nil {
		err, ok := a.value.(error)
		if !ok {
			a.tb.Fatalf("expected no error, but got different type: %T", a.value)
		}
		a.tb.Fatalf("expected no error, but got %v", err)
	}
	return a
}

func (a assertBuilder) Tokens(expected []tokenizer.Token) assertBuilder {
	a.tb.Helper()
	actual := a.value.([]tokenizer.Token)
	minLen := len(expected)
	if len(actual) < minLen {
		minLen = len(actual)
	}

	for i := 0; i < minLen; i++ {
		if expected[i].Type != actual[i].Type || expected[i].Value != actual[i].Value {
			a.tb.Fatalf("Expected %s, got %s", expected[i].String(), actual[i].String())
		}
	}

	for i := minLen; i < len(expected); i++ {
		a.tb.Logf("Expected token: %v", expected[i])
	}

	for i := minLen; i < len(actual); i++ {
		a.tb.Logf("Unexpected token: %v", actual[i])
	}

	if len(actual) != len(expected) {
		a.tb.Fatalf("Expected %d tokens, got %d", len(expected), len(actual))
	}
	return a
}

func (a assertBuilder) Program(expected ...ast.Statement) assertBuilder {
	a.tb.Helper()
	actual := a.value.(ast.Program)
	if len(actual) != len(expected) {
		a.tb.Fatalf("Expected %d statements, got %d", len(expected), len(actual))
	}
	for i, statement := range actual {
		if statement.String() != expected[i].String() {
			a.tb.Fatalf("Expected %s, got %s", expected[i].String(), statement.String())
		}
	}
	return a
}
