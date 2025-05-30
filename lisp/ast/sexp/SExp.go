package sexp

import (
	"fmt"
	"strings"

	"github.com/Olian04/go-lisp/lisp/ast"
)

type Variant string

const (
	VariantFunction Variant = "function"
	VariantOperator Variant = "operator"
	//Macro    SExpVariant = "macro"
)

type SExp struct {
	Variant    Variant
	Identifier string
	Arguments  []ast.Statement
}

func Function(identifier string, arguments []ast.Statement) SExp {
	return SExp{Variant: VariantFunction, Identifier: identifier, Arguments: arguments}
}

func Operator(identifier string, arguments []ast.Statement) SExp {
	return SExp{Variant: VariantOperator, Identifier: identifier, Arguments: arguments}
}

func (s SExp) Kind() ast.StatementKind {
	return ast.StatementKindSExp
}

func (s SExp) String() string {
	arguments := make([]string, len(s.Arguments))
	for i, argument := range s.Arguments {
		arguments[i] = argument.String()
	}
	return fmt.Sprintf("(%s %v)", s.Identifier, strings.Join(arguments, " "))
}
