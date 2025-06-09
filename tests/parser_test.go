package tests

import (
	"testing"

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
		tokens.Number("1"),
		tokens.Number("2"),
		tokens.RParen(),
	})
	util.Assert(t, err).NotError()
	util.Assert(t, program).Program(
		ast.Expression("+", ast.Number(1), ast.Number(2)),
	)
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
	util.Assert(t, program).Program(
		ast.Expression("print", ast.String("\"Hello, World!\"")),
	)
}

func TestParserNestedExpressions(t *testing.T) {
	t.Parallel()
	program, err := parser.Parse([]tokenizer.Token{
		tokens.LParen(),
		tokens.Identifier("print"),
		tokens.LParen(),
		tokens.Identifier("+"),
		tokens.Number("1"),
		tokens.Number("2"),
		tokens.Number("3"),
		tokens.RParen(),
		tokens.LParen(),
		tokens.Identifier("/"),
		tokens.Number("1"),
		tokens.Number("2"),
		tokens.RParen(),
		tokens.RParen(),
	})
	util.Assert(t, err).NotError()
	util.Assert(t, program).Program(
		ast.Expression("print",
			ast.Expression("+",
				ast.Number(1),
				ast.Number(2),
				ast.Number(3),
			),
			ast.Expression("/",
				ast.Number(1),
				ast.Number(2),
			),
		),
	)
}
