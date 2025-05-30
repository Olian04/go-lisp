package ast

import "fmt"

type StatementKind string

const (
	StatementKindSExp    StatementKind = "SExp"
	StatementKindLiteral StatementKind = "Literal"
	StatementKindInvalid StatementKind = "Invalid"
)

type Statement interface {
	Kind() StatementKind
	String() string
}

type InvalidStatement struct {
	Message string
}

func (i InvalidStatement) Kind() StatementKind {
	return StatementKindInvalid
}

func (i InvalidStatement) String() string {
	return fmt.Sprintf("Invalid: %s", i.Message)
}
