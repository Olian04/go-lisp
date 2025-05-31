package ast

import (
	"fmt"
	"strings"
)

type SExpVariant string

const (
	SExpVariantFunction SExpVariant = "function"
	SExpVariantOperator SExpVariant = "operator"
	//Macro    SExpVariant = "macro"
)

type SExp struct {
	Variant    SExpVariant
	Identifier string
	Arguments  []Statement
}

func (s SExp) Kind() StatementKind {
	return StatementKindSExp
}

func (s SExp) String() string {
	arguments := make([]string, len(s.Arguments))
	for i, argument := range s.Arguments {
		arguments[i] = argument.String()
	}
	return fmt.Sprintf("(%s %v)", s.Identifier, strings.Join(arguments, " "))
}

func Function(identifier string, arguments []Statement) SExp {
	return SExp{Variant: SExpVariantFunction, Identifier: identifier, Arguments: arguments}
}

func Operator(identifier string, arguments []Statement) SExp {
	return SExp{Variant: SExpVariantOperator, Identifier: identifier, Arguments: arguments}
}
