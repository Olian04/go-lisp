package parser

import "github.com/Olian04/go-lisp/lisp/tokenizer"

type TokenReader struct {
	tokenChan       chan tokenizer.Token
	seenTokens      []tokenizer.Token
	haveMoreTokens  bool
	committedIndex  int
	temporaryOffset int
}

func NewTokenReader(tok *tokenizer.Tokenizer) *TokenReader {
	tr := &TokenReader{
		tokenChan:       make(chan tokenizer.Token),
		seenTokens:      make([]tokenizer.Token, 0),
		haveMoreTokens:  true,
		committedIndex:  0,
		temporaryOffset: 0,
	}
	go func() {
		for {
			token := tok.NextToken()

			if token.Type == tokenizer.TokenTypeEOF || token.Type == tokenizer.TokenTypeInvalid {
				tr.tokenChan <- token
				tr.haveMoreTokens = false
				break
			}
			tr.tokenChan <- token
		}
	}()
	return tr
}

func (tr *TokenReader) Done() bool {
	return !tr.haveMoreTokens && tr.committedIndex >= len(tr.seenTokens)
}

func (tr *TokenReader) PeekToken() tokenizer.Token {
	offset := tr.temporaryOffset
	tok := tr.NextToken()
	tr.temporaryOffset = offset
	return tok
}

func (tr *TokenReader) NextToken() tokenizer.Token {
	index := tr.committedIndex + tr.temporaryOffset
	if index >= len(tr.seenTokens) {
		tok := <-tr.tokenChan
		tr.seenTokens = append(tr.seenTokens, tok)
		tr.temporaryOffset++
		return tok
	}
	token := tr.seenTokens[index]
	tr.temporaryOffset++
	return token
}

func (tr *TokenReader) Commit() {
	tr.committedIndex += tr.temporaryOffset
	tr.temporaryOffset = 0
}

func (tr *TokenReader) Rollback() {
	tr.temporaryOffset = 0
}
