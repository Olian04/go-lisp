package ast

type StatementKind string

const (
	StatementKindExpression StatementKind = "Expression"
	StatementKindLiteral    StatementKind = "Literal"
)

type Statement interface {
	Kind() StatementKind
	String() string
}
