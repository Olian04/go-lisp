package tests

import (
	"context"
	"testing"

	"github.com/Olian04/go-lisp/lisp/tokenizer"
	"github.com/Olian04/go-lisp/tests/util"
)

func TestTokenizer(t *testing.T) {
	tok := tokenizer.New(context.Background(), "(print (+ 1 2 3) (/ 1 2))")
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeIdentifier, Value: "print"})

	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeOperator, Value: "+"})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "1"})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "2"})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "3"})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"})

	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeOperator, Value: "/"})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "1"})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "2"})
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"})

	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"})
}

func TestTokenTypes(t *testing.T) {
	tok := tokenizer.New(context.Background(), "1.23")
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeFloat, Value: "1.23"})

	tok = tokenizer.New(context.Background(), "123")
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: "123"})

	tok = tokenizer.New(context.Background(), "\"1.23\"")
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeString, Value: "\"1.23\""})

	tok = tokenizer.New(context.Background(), "hello")
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeIdentifier, Value: "hello"})

	tok = tokenizer.New(context.Background(), "+")
	util.AssertToken(t, tok, tokenizer.Token{Type: tokenizer.TokenTypeOperator, Value: "+"})
}
