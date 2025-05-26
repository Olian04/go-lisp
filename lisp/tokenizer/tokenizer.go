package tokenizer

import (
	"context"
)

type Tokenizer struct {
	source    string
	readerPos int
}

func New(ctx context.Context, source string) *Tokenizer {
	return &Tokenizer{
		source:    source,
		readerPos: 0,
	}
}

func (t *Tokenizer) NextToken() Token {
	token := t.nextToken()
	if token.Type == TokenTypeEOF {
		return Token{Type: TokenTypeEOF, Value: ""}
	}
	if token.Type == tokenTypeWhitespace {
		return t.NextToken()
	}
	return token
}

func (t *Tokenizer) peek(n int) rune {
	if t.readerPos+n < len(t.source) {
		return rune(t.source[t.readerPos+n])
	}
	return 0
}

func (t *Tokenizer) read(n int) string {
	return t.source[t.readerPos : t.readerPos+n]
}

func (t *Tokenizer) incrementReader(n int) {
	t.readerPos += n
}

func (t *Tokenizer) nextToken() Token {
	if t.readerPos >= len(t.source) {
		return Token{Type: TokenTypeEOF, Value: ""}
	}
	syntax, n := t.readSyntax()
	if n > 0 {
		t.incrementReader(n)
		return syntax
	}
	number, n := t.readNumber()
	if n > 0 {
		t.incrementReader(n)
		return number
	}
	str, n := t.readString()
	if n > 0 {
		t.incrementReader(n)
		return str
	}
	identifier, n := t.readIdentifier()
	if n > 0 {
		t.incrementReader(n)
		return identifier
	}
	op, n := t.readOperator()
	if n > 0 {
		t.incrementReader(n)
		return op
	}
	whitespace, n := t.readWhitespace()
	if n > 0 {
		t.incrementReader(n)
		return whitespace
	}
	return InvalidToken
}

func (t *Tokenizer) readSyntax() (Token, int) {
	if t.peek(0) == '(' {
		return Token{Type: TokenTypeLParen, Value: t.read(1)}, 1
	} else if t.peek(0) == ')' {
		return Token{Type: TokenTypeRParen, Value: t.read(1)}, 1
	}
	return InvalidToken, 0
}

func (t *Tokenizer) readNumber() (Token, int) {
	number := t.lookahead(0, isNumberChar)
	if number > 0 {
		if t.peek(number) == '.' {
			number += 1
			number += t.lookahead(number, isNumberChar)
			return Token{Type: TokenTypeFloat, Value: t.read(number)}, number
		}
		return Token{Type: TokenTypeInteger, Value: t.read(number)}, number
	}
	return InvalidToken, 0
}

func (t *Tokenizer) readString() (Token, int) {
	if t.peek(0) == '"' {
		str := t.lookahead(1, func(r rune) bool {
			return r != '"'
		})
		str += 1

		if t.peek(str) == '"' {
			str += 1
			return Token{Type: TokenTypeString, Value: t.read(str)}, str
		}
	}
	return InvalidToken, 0
}

func (t *Tokenizer) readIdentifier() (Token, int) {
	if isIdentifierChar(t.peek(0), true) {
		identifier := t.lookahead(1, func(r rune) bool {
			return isIdentifierChar(r, false)
		})
		if identifier > 0 {
			return Token{Type: TokenTypeIdentifier, Value: t.read(identifier + 1)}, identifier + 1
		}
	}
	return InvalidToken, 0
}

func (t *Tokenizer) readOperator() (Token, int) {
	if isOperatorChar(t.peek(0)) {
		return Token{Type: TokenTypeOperator, Value: t.read(1)}, 1
	}
	return InvalidToken, 0
}

func (t *Tokenizer) readWhitespace() (Token, int) {
	whitespace := t.lookahead(0, isWhitespaceChar)
	if whitespace > 0 {
		return Token{Type: tokenTypeWhitespace, Value: t.read(whitespace)}, whitespace
	}
	return InvalidToken, 0
}

func (t *Tokenizer) lookahead(offset int, predicate func(r rune) bool) int {
	i := 0
	for predicate(t.peek(offset + i)) {
		i++
	}
	return i
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
