package parser

import (
	"context"
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

func (p *Parser) Parse() ast.Program {
	program := ast.Program{}
	tokenReader := NewTokenReader(p.tokenizer)
	program.Statements = p.parseStatements(tokenReader)
	return program
}

func (p *Parser) parseStatements(tokenReader *TokenReader) []ast.Statement {
	statements := make([]ast.Statement, 0)
	for !tokenReader.Done() {
		sexp, ok := p.parseSExp(tokenReader)
		if ok {
			statements = append(statements, sexp)
			tokenReader.Commit()
			continue
		} else {
			tokenReader.Rollback()
		}

		literal, ok := p.parseLiteral(tokenReader)
		if ok {
			statements = append(statements, literal)
			tokenReader.Commit()
			continue
		} else {
			tokenReader.Rollback()
		}
	}
	return statements
}

func (p *Parser) parseSExp(tokenReader *TokenReader) (ast.Statement, bool) {
	if tokenReader.NextToken().Type != tokenizer.TokenTypeLParen {
		return ast.InvalidStatement{Message: "Expected '('"}, false
	}
	identifier := tokenReader.NextToken()
	if identifier.Type != tokenizer.TokenTypeIdentifier && identifier.Type != tokenizer.TokenTypeOperator {
		return ast.InvalidStatement{Message: "Expected identifier or operator"}, false
	}
	if tokenReader.NextToken().Type != tokenizer.TokenTypeRParen {
		return ast.InvalidStatement{Message: "Expected ')'"}, false
	}
	return sexp.SExp{Identifier: identifier.Value}, true
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
