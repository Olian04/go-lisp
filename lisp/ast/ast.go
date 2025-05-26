package ast

import (
	"fmt"
	"strings"
)

type StatementType string

const (
	StatementTypeSExp    StatementType = "SExp"
	StatementTypeLiteral StatementType = "Literal"
)

type LiteralType string

const (
	LiteralTypeInteger LiteralType = "Integer"
	LiteralTypeFloat   LiteralType = "Float"
	LiteralTypeString  LiteralType = "String"
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

type Statement struct {
	Type    StatementType
	SExp    *SExp
	Literal *Literal
}

func (s *Statement) String() string {
	switch s.Type {
	case StatementTypeSExp:
		return s.SExp.String()
	case StatementTypeLiteral:
		return s.Literal.String()
	}
	return ""
}

type Literal struct {
	Type  LiteralType
	Value any
}

func (l *Literal) String() string {
	return fmt.Sprintf("%v", l.Value)
}

type SExp struct {
	Identifier string
	Arguments  []SExp
}

func (s *SExp) String() string {
	arguments := make([]string, len(s.Arguments))
	for i, argument := range s.Arguments {
		arguments[i] = argument.String()
	}
	return fmt.Sprintf("(%s %v)", s.Identifier, strings.Join(arguments, " "))
}
