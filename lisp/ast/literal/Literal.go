package literal

import (
	"fmt"

	"github.com/Olian04/go-lisp/lisp/ast"
)

type Variant string

const (
	VariantInteger Variant = "integer"
	VariantFloat   Variant = "float"
	VariantString  Variant = "string"
	VariantBoolean Variant = "boolean"
)

type Literal struct {
	Variant Variant
	Value   any
}

func Integer(value int) Literal {
	return Literal{Variant: VariantInteger, Value: value}
}

func Float(value float64) Literal {
	return Literal{Variant: VariantFloat, Value: value}
}

func String(value string) Literal {
	return Literal{Variant: VariantString, Value: value}
}

func Boolean(value bool) Literal {
	return Literal{Variant: VariantBoolean, Value: value}
}

func (l Literal) Kind() ast.StatementKind {
	return ast.StatementKindLiteral
}

func (l Literal) String() string {
	return fmt.Sprintf("%v", l.Value)
}
