package tests

import (
	"context"
	"testing"

	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

func TestTokenizer(t *testing.T) {
	tokenizer := tokenizer.New(context.Background(), "(+ 1 2 3)")
	tokens := []tokenizer.Token{}
	for token := range tokenizer.Tokens {
		tokens = append(tokens, token)
	}
	if len(tokens) != 5 {
		t.Errorf("Expected 5 tokens, got %d", len(tokens))
	}
	if tokens[0].Type != tokenizer.TokenTypeLParen {
		t.Errorf("Expected LParen, got %s", tokens[0])
	}
	if tokens[1].Type != tokenizer.TokenTypeIdentifier {
		t.Errorf("Expected Identifier, got %s", tokens[1])
	}
}
