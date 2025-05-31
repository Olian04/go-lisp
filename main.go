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
		tokens := tokenizer.New(input).Tokens()
		for _, token := range tokens {
			fmt.Println(token.String())
		}
		fmt.Println("--------------------------------")
		program, err := parser.New(tokens).Parse()
		if err != nil {
			fmt.Println("Error parsing program:", err)
			continue
		}
		fmt.Println(program.String())
	}
}
