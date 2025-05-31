package parser

import "github.com/Olian04/go-lisp/lisp/tokenizer"

type TokenSource interface {
	NextToken() tokenizer.Token
}
