package ast

type StatementKind string

type Statement interface {
	String() string
}
