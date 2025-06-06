package ast

import "github.com/Olian04/go-lisp/lisp/ast"

func Integer(value int) ast.Literal {
	return ast.Literal{Variant: ast.LiteralVariantInteger, Value: value}
}
func Float(value float64) ast.Literal {
	return ast.Literal{Variant: ast.LiteralVariantFloat, Value: value}
}

func String(value string) ast.Literal {
	return ast.Literal{Variant: ast.LiteralVariantString, Value: value}
}

func Boolean(value bool) ast.Literal {
	return ast.Literal{Variant: ast.LiteralVariantBoolean, Value: value}
}

func Expression(identifier string, arguments []ast.Statement) ast.Expression {
	return ast.Expression{Identifier: identifier, Arguments: arguments}
}
