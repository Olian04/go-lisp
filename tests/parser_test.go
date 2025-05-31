package tests

import (
	"testing"

	"github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/parser"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
	"github.com/Olian04/go-lisp/tests/util"
)

func TestParserSimpleExpression(t *testing.T) {
	t.Parallel()
	program, err := parser.New([]tokenizer.Token{
		tokenizer.LParen(),
		tokenizer.Operator("+"),
		tokenizer.Integer("1"),
		tokenizer.Integer("2"),
		tokenizer.RParen(),
	}).Parse()
	if err != nil {
		t.Fatalf("Expected program, got error: %s", err)
	}

	util.AssertProgram(t, program, []ast.Statement{
		ast.Operator("+", []ast.Statement{
			ast.Integer(1),
			ast.Integer(2),
		}),
	})
}

func TestParserHelloWorld(t *testing.T) {
	t.Parallel()
	program, err := parser.New([]tokenizer.Token{
		tokenizer.LParen(),
		tokenizer.Identifier("print"),
		tokenizer.String("\"Hello, World!\""),
		tokenizer.RParen(),
	}).Parse()
	if err != nil {
		t.Fatalf("Expected program, got error: %s", err)
	}

	util.AssertProgram(t, program, []ast.Statement{
		ast.Function("print", []ast.Statement{
			ast.String("\"Hello, World!\""),
		}),
	})
}

func TestParserNestedExpressions(t *testing.T) {
	t.Parallel()
	program, err := parser.New([]tokenizer.Token{
		tokenizer.LParen(),
		tokenizer.Identifier("print"),
		tokenizer.LParen(),
		tokenizer.Operator("+"),
		tokenizer.Integer("1"),
		tokenizer.Integer("2"),
		tokenizer.Integer("3"),
		tokenizer.RParen(),
		tokenizer.LParen(),
		tokenizer.Operator("/"),
		tokenizer.Integer("1"),
		tokenizer.Integer("2"),
		tokenizer.RParen(),
		tokenizer.RParen(),
	}).Parse()
	if err != nil {
		t.Fatalf("Expected program, got error: %s", err)
	}

	util.AssertProgram(t, program, []ast.Statement{
		ast.Function("print", []ast.Statement{
			ast.Function("+", []ast.Statement{
				ast.Integer(1),
				ast.Integer(2),
				ast.Integer(3),
			}),
			ast.Function("/", []ast.Statement{
				ast.Integer(1),
				ast.Integer(2),
			}),
		}),
	})
}
