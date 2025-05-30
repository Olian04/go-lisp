package tests

import (
	"testing"

	"github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/ast/literal"
	"github.com/Olian04/go-lisp/lisp/ast/sexp"
	"github.com/Olian04/go-lisp/lisp/parser"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
	"github.com/Olian04/go-lisp/tests/util"
)

func TestParserSimpleExpression(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New(t.Context(), "(+ 1 2)")
	parser := parser.New(t.Context(), tok)
	program := parser.Parse()
	util.AssertProgram(t, program, []ast.Statement{
		sexp.Operator("+", []ast.Statement{
			literal.Integer(1),
			literal.Integer(2),
		}),
	})
}

func TestParserHelloWorld(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New(t.Context(), "(print \"Hello, World!\")")
	parser := parser.New(t.Context(), tok)
	program := parser.Parse()
	util.AssertProgram(t, program, []ast.Statement{
		sexp.Function("print", []ast.Statement{
			literal.String("Hello, World!"),
		}),
	})
}

func TestParserNestedExpressions(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New(t.Context(), "(print (+ 1 2 3) (/ 1 2))")
	parser := parser.New(t.Context(), tok)
	program := parser.Parse()
	util.AssertProgram(t, program, []ast.Statement{
		sexp.Function("print", []ast.Statement{
			sexp.Function("+", []ast.Statement{
				literal.Integer(1),
				literal.Integer(2),
				literal.Integer(3),
			}),
			sexp.Function("/", []ast.Statement{
				literal.Integer(1),
				literal.Integer(2),
			}),
		}),
	})
}
