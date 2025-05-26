package tokenizer

import (
	"context"
)

type Tokenizer struct {
	source string
	pos    int
	Tokens chan Token
}

func New(ctx context.Context, source string) *Tokenizer {
	t := &Tokenizer{
		source: source,
		pos:    0,
		Tokens: make(chan Token),
	}
	go func() {
		for {
			token := t.nextToken()
			if token.Type == TokenTypeWhitespace {
				continue
			}
			if token.Type == TokenTypeEOF {
				close(t.Tokens)
				return
			}
			select {
			case <-ctx.Done():
				return
			case t.Tokens <- token:
			}
		}
	}()
	return t
}

func (t *Tokenizer) peek(offset int) rune {
	if t.pos+offset < len(t.source) {
		return rune(t.source[t.pos+offset])
	}
	return 0
}

func (t *Tokenizer) consume(n int) string {
	s := t.source[t.pos : t.pos+n]
	t.pos += n
	return s
}

func (t *Tokenizer) nextToken() Token {
	if t.pos >= len(t.source) {
		return Token{Type: TokenTypeEOF, Value: ""}
	}
	syntax, ok := t.readSyntax()
	if ok {
		return syntax
	}
	number, ok := t.readNumber()
	if ok {
		return number
	}
	str, ok := t.readString()
	if ok {
		return str
	}
	identifier, ok := t.readIdentifier()
	if ok {
		return identifier
	}
	op, ok := t.readOperator()
	if ok {
		return op
	}
	whitespace, ok := t.readWhitespace()
	if ok {
		return whitespace
	}
	return InvalidToken
}

func (t *Tokenizer) readSyntax() (Token, bool) {
	if t.peek(0) == '(' {
		return Token{Type: TokenTypeLParen, Value: t.consume(1)}, true
	} else if t.peek(0) == ')' {
		return Token{Type: TokenTypeRParen, Value: t.consume(1)}, true
	}
	return InvalidToken, false
}

func (t *Tokenizer) readNumber() (Token, bool) {
	pos := t.pos
	for isNumberChar(t.peek(0)) {
		t.consume(1)
	}
	if t.peek(0) == '.' {
		t.consume(1)
		for isNumberChar(t.peek(0)) {
			t.consume(1)
		}
		return Token{Type: TokenTypeFloat, Value: t.source[pos:t.pos]}, true
	}
	if pos != t.pos {
		return Token{Type: TokenTypeInteger, Value: t.source[pos:t.pos]}, true
	}
	return InvalidToken, false
}

func (t *Tokenizer) readString() (Token, bool) {
	if t.peek(0) == '"' {
		t.consume(1)
		for t.peek(0) != '"' {
			t.consume(1)
		}
		return Token{Type: TokenTypeString, Value: t.consume(1)}, true
	}
	return InvalidToken, false
}

func (t *Tokenizer) readIdentifier() (Token, bool) {
	pos := t.pos
	if isIdentifierChar(t.peek(0), true) {
		t.consume(1)
		for isIdentifierChar(t.peek(0), false) {
			t.consume(1)
		}
		return Token{Type: TokenTypeIdentifier, Value: t.source[pos:t.pos]}, true
	}
	return InvalidToken, false
}

func (t *Tokenizer) readOperator() (Token, bool) {
	if isOperatorChar(t.peek(0)) {
		return Token{Type: TokenTypeOperator, Value: t.consume(1)}, true
	}
	return InvalidToken, false
}

func (t *Tokenizer) readWhitespace() (Token, bool) {
	pos := t.pos
	for isWhitespaceChar(t.peek(0)) {
		t.consume(1)
	}
	if pos != t.pos {
		return Token{Type: TokenTypeWhitespace, Value: t.source[pos:t.pos]}, true
	}
	return InvalidToken, false
}

func isWhitespaceChar(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func isIdentifierChar(r rune, isFirst bool) bool {
	if isFirst {
		return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
	}
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9'
}

func isNumberChar(r rune) bool {
	return r >= '0' && r <= '9'
}

func isOperatorChar(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '%' || r == '=' || r == '<' || r == '>' || r == '&' || r == '|' || r == '^' || r == '~' || r == '!'
}
