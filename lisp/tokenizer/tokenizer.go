package tokenizer

import (
	"fmt"
)

type Tokenizer struct {
	source    string
	readerPos int
}

type tokenizerStep struct {
	token Token // token to return
	next  int   // next reader position
}

var InvalidToken = tokenizerStep{
	token: Token{Type: TokenTypeInvalid},
	next:  0,
}

var NoToken = tokenizerStep{
	token: Token{Type: TokenTypeInvalid},
	next:  0,
}

func New(source string) *Tokenizer {
	return &Tokenizer{
		source:    source,
		readerPos: 0,
	}
}

func (t *Tokenizer) Tokens() []Token {
	tokens := make([]Token, 0)
	for {
		token := t.nextToken()
		tokens = append(tokens, token)
		if token.Type == TokenTypeEOF {
			break
		}
	}
	return tokens
}

func (t *Tokenizer) nextToken() Token {
	step := t.advance()
	if step == InvalidToken {
		return Invalid(fmt.Sprintf("invalid token: %s", step.token.String()))
	}

	t.readerPos = step.next
	if step.token.Type == TokenTypeEOF {
		return EOF()
	}
	if step.token.Type == tokenTypeWhitespace {
		return t.nextToken()
	}
	return step.token
}

func (t *Tokenizer) peek(offset int) rune {
	if t.readerPos+offset < len(t.source) {
		return rune(t.source[t.readerPos+offset])
	}
	return 0
}

func (t *Tokenizer) read(size int) string {
	return t.source[t.readerPos : t.readerPos+size]
}

func (t *Tokenizer) advance() tokenizerStep {
	if t.readerPos >= len(t.source) {
		return tokenizerStep{
			token: Token{Type: TokenTypeEOF, Value: ""},
			next:  t.readerPos,
		}
	}
	if tok := t.readSyntax(); tok != NoToken {
		return tok
	}
	if tok := t.readNumber(); tok != NoToken {
		return tok
	}
	if tok := t.readString(); tok != NoToken {
		return tok
	}
	if tok := t.readIdentifier(); tok != NoToken {
		return tok
	}
	if tok := t.readOperator(); tok != NoToken {
		return tok
	}
	if tok := t.readWhitespace(); tok != NoToken {
		return tok
	}
	return InvalidToken
}

func (t *Tokenizer) readSyntax() tokenizerStep {
	if isLeftParenChar(t.peek(0)) {
		return tokenizerStep{
			token: LParen(),
			next:  t.readerPos + 1,
		}
	} else if isRightParenChar(t.peek(0)) {
		return tokenizerStep{
			token: RParen(),
			next:  t.readerPos + 1,
		}
	}
	return NoToken
}

func (t *Tokenizer) readNumber() tokenizerStep {
	number := t.lookaheadWhile(0, isNumberChar)
	if number > 0 {
		if isAccessChar(t.peek(number)) {
			number += 1
			number += t.lookaheadWhile(number, isNumberChar)
			return tokenizerStep{
				token: Float(t.read(number)),
				next:  t.readerPos + number,
			}
		}
		return tokenizerStep{
			token: Integer(t.read(number)),
			next:  t.readerPos + number,
		}
	}
	return NoToken
}

func (t *Tokenizer) readString() tokenizerStep {
	if isQuoteChar(t.peek(0)) {
		offset := 1
		offset += t.lookaheadUntil(1, isQuoteChar)

		if isQuoteChar(t.peek(offset)) {
			offset += 1
			return tokenizerStep{
				token: String(t.read(offset)),
				next:  t.readerPos + offset,
			}
		}
	}
	return NoToken
}

func (t *Tokenizer) readIdentifier() tokenizerStep {
	if isAlphaChar(t.peek(0)) {
		offset := t.lookaheadWhile(1, isIdentifierChar)
		if offset > 0 {
			return tokenizerStep{
				token: Identifier(t.read(offset + 1)),
				next:  t.readerPos + offset + 1,
			}
		}
	}
	return NoToken
}

func (t *Tokenizer) readOperator() tokenizerStep {
	if isOperatorChar(t.peek(0)) {
		return tokenizerStep{
			token: Operator(t.read(1)),
			next:  t.readerPos + 1,
		}
	}
	return NoToken
}

func (t *Tokenizer) readWhitespace() tokenizerStep {
	offset := t.lookaheadWhile(0, isWhitespaceChar)
	if offset > 0 {
		return tokenizerStep{
			token: Whitespace(t.read(offset)),
			next:  t.readerPos + offset,
		}
	}
	return NoToken
}

func (t *Tokenizer) lookaheadWhile(offset int, predicate func(r rune) bool) int {
	i := 0
	for predicate(t.peek(offset + i)) {
		i++
	}
	return i
}

func (t *Tokenizer) lookaheadUntil(offset int, predicate func(r rune) bool) int {
	i := 0
	for !predicate(t.peek(offset + i)) {
		i++
	}
	return i
}

func isWhitespaceChar(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func isIdentifierChar(r rune) bool {
	return isAlphaChar(r) || isNumberChar(r) || isSeparatorChar(r)
}

func isQuoteChar(r rune) bool {
	return r == '"'
}

func isLeftParenChar(r rune) bool {
	return r == '('
}

func isRightParenChar(r rune) bool {
	return r == ')'
}

func isLeftBraceChar(r rune) bool {
	return r == '{'
}

func isRightBraceChar(r rune) bool {
	return r == '}'
}

func isLeftBracketChar(r rune) bool {
	return r == '['
}

func isRightBracketChar(r rune) bool {
	return r == ']'
}

func isAccessChar(r rune) bool {
	return r == '.'
}

func isSeparatorChar(r rune) bool {
	return r == '_'
}

func isAlphaChar(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isNumberChar(r rune) bool {
	return r >= '0' && r <= '9'
}

func isOperatorChar(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '%' || r == '=' || r == '<' || r == '>' || r == '&' || r == '|' || r == '^' || r == '~' || r == '!'
}
