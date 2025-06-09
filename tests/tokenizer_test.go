package tests

import (
	"testing"

	"github.com/Olian04/go-lisp/lisp/tokenizer"
	"github.com/Olian04/go-lisp/tests/util"
	"github.com/Olian04/go-lisp/tests/util/tokens"
)

func TestTokenizerEOF(t *testing.T) {
	t.Parallel()
	tok, err := tokenizer.Tokenize("")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
		tokens.EOF(),
	})
}

func TestExpression(t *testing.T) {
	t.Parallel()
	tok, err := tokenizer.Tokenize("(+ 1 2)")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
		tokens.LParen(),
		tokens.Identifier("+"),
		tokens.Number("1"),
		tokens.Number("2"),
		tokens.RParen(),
		tokens.EOF(),
	})
}

func TestHelloWorld(t *testing.T) {
	t.Parallel()
	tok, err := tokenizer.Tokenize("(print \"Hello, World!\")")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
		tokens.LParen(),
		tokens.Identifier("print"),
		tokens.String("\"Hello, World!\""),
		tokens.RParen(),
		tokens.EOF(),
	})
}

func TestNestedExpressions(t *testing.T) {
	t.Parallel()
	tok, err := tokenizer.Tokenize("(print (+ 1 2 3) (/ 1 2))")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
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
		tokens.EOF(),
	})
}

func TestTokenTypes(t *testing.T) {
	t.Parallel()
	tok, err := tokenizer.Tokenize("1.23")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
		tokens.Number("1.23"),
		tokens.EOF(),
	})

	tok, err = tokenizer.Tokenize("123")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
		tokens.Number("123"),
		tokens.EOF(),
	})

	tok, err = tokenizer.Tokenize("\"1.23\"")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
		tokens.String("\"1.23\""),
		tokens.EOF(),
	})

	tok, err = tokenizer.Tokenize("hello")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
		tokens.Identifier("hello"),
		tokens.EOF(),
	})

	tok, err = tokenizer.Tokenize("+")
	util.Assert(t, err).NotError()
	util.Assert(t, tok).Tokens([]tokenizer.Token{
		tokens.Identifier("+"),
		tokens.EOF(),
	})
}
