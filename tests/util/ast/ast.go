package ast

import "github.com/Olian04/go-lisp/lisp/parser/ast"

func Number(value float64) ast.Literal {
	return ast.Literal{Variant: ast.LiteralVariantNumber, Value: value}
}

func String(value string) ast.Literal {
	return ast.Literal{Variant: ast.LiteralVariantString, Value: value}
}

func Boolean(value bool) ast.Literal {
	return ast.Literal{Variant: ast.LiteralVariantBoolean, Value: value}
}

func Expression(identifier string, arguments ...ast.Statement) ast.Expression {
	return ast.Expression{Identifier: identifier, Arguments: arguments}
}

func Program(statements ...ast.Statement) ast.Program {
	return statements
}
