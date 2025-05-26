package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	tokenizer := tokenizer.New(ctx, "(+ 1 2 3)")
	for token := range tokenizer.Tokens {
		fmt.Println(token)
	}
}
