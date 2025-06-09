package builtins

import (
	"fmt"
	"strings"

	"github.com/Olian04/go-lisp/lisp/evaluator/context"
	"github.com/Olian04/go-lisp/lisp/evaluator/util"
	"github.com/Olian04/go-lisp/lisp/parser/ast"
)

func Print(args []ast.Statement, context context.EvaluatorContext) {
	util.AssertAtLeastArgs(args, 1)
	strs := make([]string, len(args))
	for i, arg := range args {
		strs[i] = arg.String()
	}
	fmt.Fprint(context.StdOut, strings.Join(strs, " "))
}
