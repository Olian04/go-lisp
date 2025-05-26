package parser

import (
	"context"
	"fmt"

	"github.com/Olian04/go-lisp/lisp/ast"
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
	for {
		token := p.tokenizer.NextToken()
		if token.Type == tokenizer.TokenTypeEOF {
			break
		}
		program.Statements = append(program.Statements, p.parseStatement())
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	token := p.tokenizer.NextToken()
	switch token.Type {
	case tokenizer.TokenTypeLParen:
		return p.parseSExp()
	case tokenizer.TokenTypeIdentifier:
		return p.parseLiteral()
	case tokenizer.TokenTypeEOF:
		return ast.Statement{}
	}
	panic(fmt.Sprintf("Unexpected token: %s", token.String()))
}

func (p *Parser) parseSExp() ast.Statement {
	token := p.tokenizer.NextToken()
	if token.Type == tokenizer.TokenTypeLParen {
		return ast.Statement{
			Type: ast.StatementTypeSExp,
			SExp: &ast.SExp{
				Identifier: token.Value,
			},
		}
	}
	panic(fmt.Sprintf("Unexpected token: %s", token.String()))
}

func (p *Parser) parseLiteral() ast.Statement {
	token := p.tokenizer.NextToken()
	return ast.Statement{
		Type: ast.StatementTypeLiteral,
		Literal: &ast.Literal{
			Type:  ast.LiteralTypeInteger,
			Value: token.Value,
		},
	}
}
