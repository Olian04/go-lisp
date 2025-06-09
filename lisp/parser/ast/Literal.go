package ast

import (
	"fmt"
)

type LiteralVariant string

const (
	LiteralVariantNumber  LiteralVariant = "number"
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
	case LiteralVariantNumber:
		return fmt.Sprintf("%f", l.Value)
	case LiteralVariantString:
		return fmt.Sprintf("%s", l.Value)
	case LiteralVariantBoolean:
		return fmt.Sprintf("%t", l.Value)
	default:
		return fmt.Sprintf("%v", l.Value)
	}
}
