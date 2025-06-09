package tests

import (
	"bytes"
	"testing"

	"github.com/Olian04/go-lisp/lisp/evaluator"
	"github.com/Olian04/go-lisp/lisp/evaluator/context"
	"github.com/Olian04/go-lisp/tests/util"
	"github.com/Olian04/go-lisp/tests/util/ast"
)

func TestEvaluatorPrint(t *testing.T) {
	t.Parallel()
	program := ast.Program(
		ast.Expression("print", ast.String("Hello, World!")),
	)

	var buf bytes.Buffer
	err := evaluator.Evaluate(program, context.EvaluatorContext{
		StdOut: &buf,
	})
	util.Assert(t, err).NotError()
	util.Assert(t, buf.String()).Equal("Hello, World!")
}

func TestEvaluatorAdd(t *testing.T) {
	t.Parallel()
	program := ast.Program(
		ast.Expression("print",
			ast.Expression("+", ast.Number(1), ast.Number(2)),
		),
	)

	var buf bytes.Buffer
	err := evaluator.Evaluate(program, context.EvaluatorContext{
		StdOut: &buf,
	})
	util.Assert(t, err).NotError()
	util.Assert(t, buf.String()).Equal("3")
}

func TestEvaluatorSub(t *testing.T) {
	t.Parallel()
	program := ast.Program(
		ast.Expression("print",
			ast.Expression("-", ast.Number(1), ast.Number(2)),
		),
	)

	var buf bytes.Buffer
	err := evaluator.Evaluate(program, context.EvaluatorContext{
		StdOut: &buf,
	})
	util.Assert(t, err).NotError()
	util.Assert(t, buf.String()).Equal("-1")
}
