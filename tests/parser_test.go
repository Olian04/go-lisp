package tests

import (
	"testing"

	internalAst "github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/parser"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
	"github.com/Olian04/go-lisp/tests/util"
	"github.com/Olian04/go-lisp/tests/util/ast"
	"github.com/Olian04/go-lisp/tests/util/tokens"
)

func TestParserSimpleExpression(t *testing.T) {
	t.Parallel()
	program, err := parser.Parse([]tokenizer.Token{
		tokens.LParen(),
		tokens.Identifier("+"),
		tokens.Integer("1"),
		tokens.Integer("2"),
		tokens.RParen(),
	})
	util.Assert(t, err).NotError()
	util.Assert(t, program).Program([]internalAst.Statement{
		ast.Expression("+", []internalAst.Statement{
			ast.Integer(1),
			ast.Integer(2),
		}),
	})
}

func TestParserHelloWorld(t *testing.T) {
	t.Parallel()
	program, err := parser.Parse([]tokenizer.Token{
		tokens.LParen(),
		tokens.Identifier("print"),
		tokens.String("\"Hello, World!\""),
		tokens.RParen(),
	})
	util.Assert(t, err).NotError()
	util.Assert(t, program).Program([]internalAst.Statement{
		ast.Expression("print", []internalAst.Statement{
			ast.String("\"Hello, World!\""),
		}),
	})
}

func TestParserNestedExpressions(t *testing.T) {
	t.Parallel()
	program, err := parser.Parse([]tokenizer.Token{
		tokens.LParen(),
		tokens.Identifier("print"),
		tokens.LParen(),
		tokens.Identifier("+"),
		tokens.Integer("1"),
		tokens.Integer("2"),
		tokens.Integer("3"),
		tokens.RParen(),
		tokens.LParen(),
		tokens.Identifier("/"),
		tokens.Integer("1"),
		tokens.Integer("2"),
		tokens.RParen(),
		tokens.RParen(),
	})
	util.Assert(t, err).NotError()
	util.Assert(t, program).Program([]internalAst.Statement{
		ast.Expression("print", []internalAst.Statement{
			ast.Expression("+", []internalAst.Statement{
				ast.Integer(1),
				ast.Integer(2),
				ast.Integer(3),
			}),
			ast.Expression("/", []internalAst.Statement{
				ast.Integer(1),
				ast.Integer(2),
			}),
		}),
	})
}
