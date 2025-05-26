package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Olian04/go-lisp/lisp/parser"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	tok := tokenizer.New(ctx, "(+ 1 2 3)")
	for {
		token := tok.NextToken()
		if token.Type == tokenizer.TokenTypeEOF {
			break
		}
		fmt.Println(token)
	}
	parser := parser.New(ctx, tok)
	program := parser.Parse()
	fmt.Println(program.String())
}
