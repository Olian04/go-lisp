package util

import (
	"testing"

	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

func AssertToken(t *testing.T, tok *tokenizer.Tokenizer, token tokenizer.Token) {
	actual := tok.NextToken()
	if actual.Type != token.Type || actual.Value != token.Value {
		t.Fatalf("Expected %s, got %s", token.String(), actual.String())
	}
}
