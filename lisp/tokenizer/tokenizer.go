package tokenizer

import (
	"fmt"
)

type tokenizerState struct {
	source    string
	readerPos int
	token     Token
}

func InvalidToken(source string) tokenizerState {
	return tokenizerState{
		source:    source,
		readerPos: 0,
		token:     Token{Type: TokenTypeInvalid},
	}
}

var Nothing = tokenizerState{
	source:    "",
	readerPos: 0,
	token:     Token{Type: TokenTypeNothing},
}

var EndOfFile = tokenizerState{
	source:    "",
	readerPos: 0,
	token:     Token{Type: TokenTypeEOF},
}

func Tokenize(source string) ([]Token, error) {
	state := tokenizerState{
		source:    source,
		readerPos: 0,
	}

	tokens := make([]Token, 0)
	for {
		state = nextToken(state)
		tokens = append(tokens, state.token)
		if state.token.Type == TokenTypeInvalid {
			return nil, fmt.Errorf("invalid token: %s", state.token.String())
		}
		if state.token.Type == TokenTypeEOF {
			break
		}
	}
	return tokens, nil
}

func nextToken(state tokenizerState) tokenizerState {
	step := readToken(state)
	if step.token.Type == TokenTypeInvalid {
		return InvalidToken(fmt.Sprintf("invalid token: %s", step.token.String()))
	}
	if step.token.Type == TokenTypeEOF {
		return EndOfFile
	}
	if step.token.Type == TokenTypeWhitespace {
		return nextToken(step)
	}
	return step
}

func peek(state tokenizerState, offset int) rune {
	if state.readerPos+offset < len(state.source) {
		return rune(state.source[state.readerPos+offset])
	}
	return 0
}

func read(state tokenizerState, size int) string {
	return state.source[state.readerPos : state.readerPos+size]
}

func readToken(state tokenizerState) tokenizerState {
	if state.readerPos >= len(state.source) {
		return EndOfFile
	}
	if tok := readSyntax(state); tok.token.Type != TokenTypeNothing {
		return tok
	}
	if tok := readNumber(state); tok.token.Type != TokenTypeNothing {
		return tok
	}
	if tok := readString(state); tok.token.Type != TokenTypeNothing {
		return tok
	}
	if tok := readIdentifier(state); tok.token.Type != TokenTypeNothing {
		return tok
	}
	if tok := readOperator(state); tok.token.Type != TokenTypeNothing {
		return tok
	}
	if tok := readWhitespace(state); tok.token.Type != TokenTypeNothing {
		return tok
	}
	return InvalidToken(fmt.Sprintf("invalid token at %d: %s", state.readerPos, state.source[state.readerPos:min(state.readerPos+10, len(state.source))]))
}

func readSyntax(state tokenizerState) tokenizerState {
	if isLeftParenChar(peek(state, 0)) {
		return tokenizerState{
			source:    state.source,
			token:     Token{Type: TokenTypeLParen, Value: read(state, 1)},
			readerPos: state.readerPos + 1,
		}
	} else if isRightParenChar(peek(state, 0)) {
		return tokenizerState{
			source:    state.source,
			token:     Token{Type: TokenTypeRParen, Value: read(state, 1)},
			readerPos: state.readerPos + 1,
		}
	}
	return Nothing
}

func readNumber(state tokenizerState) tokenizerState {
	number := lookaheadWhile(state, 0, isNumberChar)
	if number > 0 {
		if isFloatSeparatorChar(peek(state, number)) {
			number += 1
			number += lookaheadWhile(state, number, isNumberChar)
			return tokenizerState{
				source:    state.source,
				token:     Token{Type: TokenTypeFloat, Value: read(state, number)},
				readerPos: state.readerPos + number,
			}
		}
		return tokenizerState{
			source:    state.source,
			token:     Token{Type: TokenTypeInteger, Value: read(state, number)},
			readerPos: state.readerPos + number,
		}
	}
	return Nothing
}

func readString(state tokenizerState) tokenizerState {
	if isQuoteChar(peek(state, 0)) {
		offset := 1
		offset += lookaheadUntil(state, 1, isQuoteChar)

		if isQuoteChar(peek(state, offset)) {
			offset += 1
			return tokenizerState{
				source:    state.source,
				token:     Token{Type: TokenTypeString, Value: read(state, offset)},
				readerPos: state.readerPos + offset,
			}
		}
	}
	return Nothing
}

func readIdentifier(state tokenizerState) tokenizerState {
	if isAlphaChar(peek(state, 0)) {
		offset := lookaheadWhile(state, 1, isIdentifierChar)
		if offset > 0 {
			return tokenizerState{
				source:    state.source,
				token:     Token{Type: TokenTypeIdentifier, Value: read(state, offset+1)},
				readerPos: state.readerPos + offset + 1,
			}
		}
	}
	return Nothing
}

func readOperator(state tokenizerState) tokenizerState {
	offset := lookaheadWhile(state, 0, isOperatorChar)
	if offset > 0 {
		return tokenizerState{
			source:    state.source,
			token:     Token{Type: TokenTypeIdentifier, Value: read(state, offset)},
			readerPos: state.readerPos + offset,
		}
	}
	return Nothing
}

func readWhitespace(state tokenizerState) tokenizerState {
	offset := lookaheadWhile(state, 0, isWhitespaceChar)
	if offset > 0 {
		return tokenizerState{
			source:    state.source,
			token:     Token{Type: TokenTypeWhitespace, Value: read(state, offset)},
			readerPos: state.readerPos + offset,
		}
	}
	return Nothing
}

func lookaheadWhile(state tokenizerState, offset int, predicate func(r rune) bool) int {
	i := 0
	for predicate(peek(state, offset+i)) {
		i++
	}
	return i
}

func lookaheadUntil(state tokenizerState, offset int, predicate func(r rune) bool) int {
	i := 0
	for !predicate(peek(state, offset+i)) {
		i++
	}
	return i
}
