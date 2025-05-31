package tests

import (
	"testing"

	"github.com/Olian04/go-lisp/lisp/tokenizer"
	"github.com/Olian04/go-lisp/tests/util"
)

func TestTokenizerEOF(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New("")
	util.AssertTokens(t, tok, []tokenizer.Token{
		tokenizer.EOF(),
	})
}

func TestExpression(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New("(+ 1 2)")
	util.AssertTokens(t, tok, []tokenizer.Token{
		tokenizer.LParen(),
		tokenizer.Operator("+"),
		tokenizer.Integer("1"),
		tokenizer.Integer("2"),
		tokenizer.RParen(),
	})
}

func TestHelloWorld(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New("(print \"Hello, World!\")")
	util.AssertTokens(t, tok, []tokenizer.Token{
		tokenizer.LParen(),
		tokenizer.Identifier("print"),
		tokenizer.String("\"Hello, World!\""),
		tokenizer.RParen(),
	})
}

func TestNestedExpressions(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New("(print (+ 1 2 3) (/ 1 2))")
	util.AssertTokens(t, tok, []tokenizer.Token{
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
	})
}

func TestTokenTypes(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New("1.23")
	util.AssertTokens(t, tok, []tokenizer.Token{
		tokenizer.Float("1.23"),
	})

	tok = tokenizer.New("123")
	util.AssertTokens(t, tok, []tokenizer.Token{
		tokenizer.Integer("123"),
	})

	tok = tokenizer.New("\"1.23\"")
	util.AssertTokens(t, tok, []tokenizer.Token{
		tokenizer.String("\"1.23\""),
	})

	tok = tokenizer.New("hello")
	util.AssertTokens(t, tok, []tokenizer.Token{
		tokenizer.Identifier("hello"),
	})

	tok = tokenizer.New("+")
	util.AssertTokens(t, tok, []tokenizer.Token{
		tokenizer.Operator("+"),
	})
}
