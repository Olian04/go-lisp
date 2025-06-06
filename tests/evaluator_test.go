package tests

import (
	"bytes"
	"testing"

	"github.com/Olian04/go-lisp/lisp/evaluator"
	"github.com/Olian04/go-lisp/lisp/evaluator/context"
	"github.com/Olian04/go-lisp/lisp/parser"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
	"github.com/Olian04/go-lisp/tests/util"
	"github.com/Olian04/go-lisp/tests/util/tokens"
)

func TestEvaluatorPrint(t *testing.T) {
	t.Parallel()
	program, err := parser.Parse([]tokenizer.Token{
		tokens.LParen(),
		tokens.Identifier("print"),
		tokens.String("Hello, World!"),
		tokens.RParen(),
	})
	util.Assert(t, err).NotError()

	var buf bytes.Buffer
	err = evaluator.Evaluate(program, context.EvaluatorContext{
		StdOut: &buf,
	})
	util.Assert(t, err).NotError()
	util.Assert(t, buf.String()).Equal("Hello, World!")
}
