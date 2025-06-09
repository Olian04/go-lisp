package parser

import (
	"fmt"
	"strconv"

	"github.com/Olian04/go-lisp/lisp/parser/ast"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

func Parse(tokens []tokenizer.Token) (ast.Program, error) {
	statements, remaining := parseStatements(tokens)
	if len(remaining) > 0 {
		return ast.Program{}, fmt.Errorf("unexpected tokens after end of program: %v", remaining)
	}
	return statements, nil
}

func parseStatements(tokens []tokenizer.Token) ([]ast.Statement, []tokenizer.Token) {
	statements := make([]ast.Statement, 0)
	remaining := tokens
	maxIterations := len(tokens) * 1000
	for len(remaining) > 0 {
		var stmt ast.Statement
		stmt, remaining = parseStatement(remaining)
		if stmt != nil {
			statements = append(statements, stmt)
		}
		if maxIterations == 0 {
			panic("Reached max iterations, possible infinite loop")
		}
		maxIterations--
	}
	return statements, remaining
}

func parseStatement(tokens []tokenizer.Token) (ast.Statement, []tokenizer.Token) {
	sexp, remaining := parseSExp(tokens)
	if sexp != nil {
		return sexp, remaining
	}

	literal, remaining := parseLiteral(tokens)
	if literal != nil {
		return literal, remaining
	}
	return nil, remaining
}

func parseSExp(tokens []tokenizer.Token) (ast.Statement, []tokenizer.Token) {
	tok, remaining := nextToken(tokens)
	if tok.Type != tokenizer.TokenTypeLParen {
		return nil, tokens
	}
	identifier, remaining := nextToken(remaining)
	if identifier.Type != tokenizer.TokenTypeIdentifier {
		return nil, tokens
	}
	arguments := make([]ast.Statement, 0)
	for {
		tok = peekToken(remaining)
		if tok.Type == tokenizer.TokenTypeRParen {
			_, remaining = nextToken(remaining)
			break
		}
		var stmt ast.Statement
		stmt, remaining = parseStatement(remaining)
		if stmt == nil {
			return nil, tokens
		}
		arguments = append(arguments, stmt)
	}
	if tok.Type != tokenizer.TokenTypeRParen {
		return nil, tokens
	}
	return ast.Expression{
		Identifier: identifier.Value,
		Arguments:  arguments,
	}, remaining
}

func parseLiteral(tokens []tokenizer.Token) (ast.Statement, []tokenizer.Token) {
	tok, remaining := nextToken(tokens)
	switch tok.Type {
	case tokenizer.TokenTypeNumber:
		value, err := strconv.ParseFloat(tok.Value, 64)
		if err != nil {
			return nil, tokens
		}
		return ast.Literal{
			Variant: ast.LiteralVariantNumber,
			Value:   value,
		}, remaining
	case tokenizer.TokenTypeString:
		return ast.Literal{
			Variant: ast.LiteralVariantString,
			Value:   tok.Value,
		}, remaining
	default:
		return nil, tokens
	}
}

func nextToken(tokens []tokenizer.Token) (tokenizer.Token, []tokenizer.Token) {
	if len(tokens) == 0 {
		return tokenizer.Token{Type: tokenizer.TokenTypeEOF}, tokens
	}
	return tokens[0], tokens[1:]
}

func peekToken(tokens []tokenizer.Token) tokenizer.Token {
	if len(tokens) == 0 {
		return tokenizer.Token{Type: tokenizer.TokenTypeEOF}
	}
	return tokens[0]
}
