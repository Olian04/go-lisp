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
	if t.peek(0) == '(' {
		return tokenizerStep{
			token: LParen(),
			next:  t.readerPos + 1,
		}
	} else if t.peek(0) == ')' {
		return tokenizerStep{
			token: RParen(),
			next:  t.readerPos + 1,
		}
	}
	return NoToken
}

func (t *Tokenizer) readNumber() tokenizerStep {
	number := t.lookahead(0, isNumberChar)
	if number > 0 {
		if t.peek(number) == '.' {
			number += 1
			number += t.lookahead(number, isNumberChar)
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
	if t.peek(0) == '"' {
		str := t.lookahead(1, func(r rune) bool {
			return r != '"'
		})
		str += 1

		if t.peek(str) == '"' {
			str += 1
			return tokenizerStep{
				token: String(t.read(str)),
				next:  t.readerPos + str,
			}
		}
	}
	return NoToken
}

func (t *Tokenizer) readIdentifier() tokenizerStep {
	if isIdentifierChar(t.peek(0), true) {
		identifier := t.lookahead(1, func(r rune) bool {
			return isIdentifierChar(r, false)
		})
		if identifier > 0 {
			return tokenizerStep{
				token: Identifier(t.read(identifier + 1)),
				next:  t.readerPos + identifier + 1,
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
	whitespace := t.lookahead(0, isWhitespaceChar)
	if whitespace > 0 {
		return tokenizerStep{
			token: Whitespace(t.read(whitespace)),
			next:  t.readerPos + whitespace,
		}
	}
	return NoToken
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
