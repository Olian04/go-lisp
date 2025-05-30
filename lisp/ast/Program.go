package ast

import (
	"strings"
)

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	statements := make([]string, len(p.Statements))
	for i, statement := range p.Statements {
		statements[i] = statement.String()
	}
	return strings.Join(statements, "\n")
}
