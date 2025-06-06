package ast

import (
	"strings"
)

type Program []Statement

func (p Program) String() string {
	statements := make([]string, len(p))
	for i, statement := range p {
		statements[i] = statement.String()
	}
	return strings.Join(statements, "\n")
}
