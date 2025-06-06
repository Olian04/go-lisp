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
	switch l.Variant {
	case LiteralVariantInteger:
		return fmt.Sprintf("%d", l.Value)
	case LiteralVariantFloat:
		return fmt.Sprintf("%f", l.Value)
	case LiteralVariantString:
		return fmt.Sprintf("\"%s\"", l.Value)
	case LiteralVariantBoolean:
		return fmt.Sprintf("%t", l.Value)
	default:
		return fmt.Sprintf("%v", l.Value)
	}
}
