package ast

type StatementKind string

const (
	StatementKindExpression StatementKind = "Expression"
	StatementKindLiteral    StatementKind = "Literal"
	StatementKindNothing    StatementKind = "Nothing"
)

type Statement interface {
	Kind() StatementKind
	String() string
}
