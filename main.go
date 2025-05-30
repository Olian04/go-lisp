package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Olian04/go-lisp/lisp/parser"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}
		tok := tokenizer.New(context.Background(), input)
		for token := tok.NextToken(); token.Type != tokenizer.TokenTypeEOF; token = tok.NextToken() {
			fmt.Println(token.String())
		}
		tok = tokenizer.New(context.Background(), input) // reset the tokenizer
		fmt.Println("--------------------------------")
		parser := parser.New(context.Background(), tok)
		program := parser.Parse()
		fmt.Println(program.String())
	}
}
