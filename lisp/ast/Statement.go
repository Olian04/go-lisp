package ast

type StatementKind string

const (
	StatementKindSExp    StatementKind = "SExp"
	StatementKindLiteral StatementKind = "Literal"
	StatementKindNothing StatementKind = "Nothing"
)

type Statement interface {
	Kind() StatementKind
	String() string
}
