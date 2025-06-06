package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Olian04/go-lisp/lisp/evaluator"
	"github.com/Olian04/go-lisp/lisp/evaluator/context"
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

		tokens, err := tokenizer.Tokenize(input)
		if err != nil {
			fmt.Println("Error tokenizing input:", err)
			continue
		}
		program, err := parser.Parse(tokens)
		if err != nil {
			fmt.Println("Error parsing program:", err)
			continue
		}

		err = evaluator.Evaluate(program, context.EvaluatorContext{
			StdOut: os.Stdout,
		})
		if err != nil {
			fmt.Println("Error evaluating program:", err)
		}
	}
}
