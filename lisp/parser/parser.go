package parser

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/ast/literal"
	"github.com/Olian04/go-lisp/lisp/ast/sexp"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

type Parser struct {
	tokenizer *tokenizer.Tokenizer
}

func New(ctx context.Context, tokenizer *tokenizer.Tokenizer) *Parser {
	return &Parser{tokenizer: tokenizer}
}

func (p *Parser) Parse() (ast.Program, bool) {
	program := ast.Program{}
	tokenReader := NewTokenReader(p.tokenizer)
	statements, ok := p.parseStatements(tokenReader)
	program.Statements = statements
	if !ok {
		return program, false
	}
	return program, true
}

func (p *Parser) parseStatements(tokenReader *TokenReader) ([]ast.Statement, bool) {
	statements := make([]ast.Statement, 0)
	for !tokenReader.Done() {
		stmt := p.parseStatement(tokenReader)
		statements = append(statements, stmt)
		if stmt.Kind() == ast.StatementKindInvalid {
			return statements, false
		}
	}
	return statements, true
}

func (p *Parser) parseStatement(tokenReader *TokenReader) ast.Statement {
	sexp, ok := p.parseSExp(tokenReader)
	if ok {
		tokenReader.Commit()
		return sexp
	} else {
		tokenReader.Rollback()
	}

	literal, ok := p.parseLiteral(tokenReader)
	if ok {
		tokenReader.Commit()
		return literal
	} else {
		tokenReader.Rollback()
	}
	return ast.InvalidStatement{Message: "Expected statement"}
}

func (p *Parser) parseSExp(tokenReader *TokenReader) (ast.Statement, bool) {
	if tok := tokenReader.NextToken(); tok.Type != tokenizer.TokenTypeLParen {
		return ast.InvalidStatement{Message: fmt.Sprintf("Expected '(' but got %s", tok.String())}, false
	}
	identifier := tokenReader.NextToken()
	if identifier.Type != tokenizer.TokenTypeIdentifier && identifier.Type != tokenizer.TokenTypeOperator {
		return ast.InvalidStatement{Message: fmt.Sprintf("Expected identifier or operator but got %s", identifier.String())}, false
	}
	arguments := make([]ast.Statement, 0)
	for {
		tok := tokenReader.NextToken()
		if tok.Type == tokenizer.TokenTypeRParen {
			break
		}
		arguments = append(arguments, p.parseStatement(tokenReader))
	}
	if tok := tokenReader.NextToken(); tok.Type != tokenizer.TokenTypeRParen {
		return ast.InvalidStatement{Message: fmt.Sprintf("Expected ')' but got %s", tok.String())}, false
	}
	return sexp.SExp{Identifier: identifier.Value, Arguments: arguments}, true
}

func (p *Parser) parseLiteral(tokenReader *TokenReader) (ast.Statement, bool) {
	token := tokenReader.NextToken()
	switch token.Type {
	case tokenizer.TokenTypeInteger:
		value, err := strconv.ParseInt(token.Value, 10, 64)
		if err != nil {
			return ast.InvalidStatement{Message: "Invalid integer literal"}, false
		}
		return literal.Integer(int(value)), true
	case tokenizer.TokenTypeFloat:
		value, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return ast.InvalidStatement{Message: "Invalid float literal"}, false
		}
		return literal.Float(value), true
	case tokenizer.TokenTypeString:
		return literal.String(token.Value), true
	default:
		return ast.InvalidStatement{Message: "Expected integer, float, or string"}, false
	}
}
