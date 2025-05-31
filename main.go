package main

import (
	"bufio"
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
		tok := tokenizer.New(input)
		tokens := tok.Array()
		for _, token := range tokens {
			fmt.Println(token.String())
		}
		fmt.Println("--------------------------------")
		parser := parser.New(tokens)
		program, err := parser.Parse()
		if err != nil {
			fmt.Println("Error parsing program:", err)
			continue
		}
		fmt.Println(program.String())
	}
}
