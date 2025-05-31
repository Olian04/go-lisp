package ast

import (
	"fmt"
)

type LiteralVariant string

const (
	LiteralVariantInteger LiteralVariant = "integer"
	LiteralVariantFloat   LiteralVariant = "float"
	LiteralVariantString  LiteralVariant = "string"
	LiteralVariantBoolean LiteralVariant = "boolean"
)

type Literal struct {
	Variant LiteralVariant
	Value   any
}

func (l Literal) Kind() StatementKind {
	return StatementKindLiteral
}

func (l Literal) String() string {
	return fmt.Sprintf("%v", l.Value)
}

func Integer(value int) Literal {
	return Literal{Variant: LiteralVariantInteger, Value: value}
}

func Float(value float64) Literal {
	return Literal{Variant: LiteralVariantFloat, Value: value}
}

func String(value string) Literal {
	return Literal{Variant: LiteralVariantString, Value: value}
}

func Boolean(value bool) Literal {
	return Literal{Variant: LiteralVariantBoolean, Value: value}
}
