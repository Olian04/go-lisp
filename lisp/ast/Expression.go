package ast

import (
	"fmt"
	"strings"
)

type Expression struct {
	Identifier string
	Arguments  []Statement
}

func (s Expression) Kind() StatementKind {
	return StatementKindExpression
}

func (s Expression) String() string {
	arguments := make([]string, len(s.Arguments))
	for i, argument := range s.Arguments {
		arguments[i] = argument.String()
	}
	return fmt.Sprintf("(%s %v)", s.Identifier, strings.Join(arguments, " "))
}
