package util

import (
	"github.com/Olian04/go-lisp/lisp/parser"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

type mockTokenSource struct {
	tokens []tokenizer.Token
	index  int
}

func (m *mockTokenSource) NextToken() tokenizer.Token {
	if m.index >= len(m.tokens) {
		return tokenizer.Token{Type: tokenizer.TokenTypeEOF}
	}
	token := m.tokens[m.index]
	m.index++
	return token
}

func NewMockTokenSource(tokens []tokenizer.Token) parser.TokenSource {
	return &mockTokenSource{tokens: tokens, index: 0}
}
