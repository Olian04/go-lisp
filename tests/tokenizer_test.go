package tests

import (
	"context"
	"testing"

	"github.com/Olian04/go-lisp/lisp/tokenizer"
	"github.com/Olian04/go-lisp/tests/util"
)

func TestTokenizerEOF(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New(context.Background(), "")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeEOF})
}

func TestExpression(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New(context.Background(), "(+ 1 2)")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeOperator, Value: "+"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "1"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "2"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"})
}

func TestHelloWorld(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New(context.Background(), "(print \"Hello, World!\")")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeIdentifier, Value: "print"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeString, Value: "\"Hello, World!\""})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"})
}

func TestNestedExpressions(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New(context.Background(), "(print (+ 1 2 3) (/ 1 2))")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeIdentifier, Value: "print"})

	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeOperator, Value: "+"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "1"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "2"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "3"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"})

	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeOperator, Value: "/"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "1"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "2"})
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"})

	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"})
}

func TestTokenTypes(t *testing.T) {
	t.Parallel()
	tok := tokenizer.New(context.Background(), "1.23")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeFloat, Value: "1.23"})

	tok = tokenizer.New(context.Background(), "123")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "123"})

	tok = tokenizer.New(context.Background(), "\"1.23\"")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeString, Value: "\"1.23\""})

	tok = tokenizer.New(context.Background(), "hello")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeIdentifier, Value: "hello"})

	tok = tokenizer.New(context.Background(), "+")
	util.AssertNextToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeOperator, Value: "+"})
}
