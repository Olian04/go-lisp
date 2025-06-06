package builtins

import (
	"fmt"
	"strings"

	"github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/evaluator/context"
)

func Print(args []ast.Statement, context context.EvaluatorContext) {
	strs := make([]string, len(args))
	for i, arg := range args {
		strs[i] = arg.String()
	}
	fmt.Fprintln(context.StdOut, strings.Join(strs, " "))
}
