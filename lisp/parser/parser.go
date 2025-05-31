package parser

import (
	"fmt"
	"strconv"

	"github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/tokenizer"
)

type Parser struct {
	tokens []tokenizer.Token
}

func New(tokens []tokenizer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() (ast.Program, error) {
	statements, remaining := p.parseStatements(p.tokens)
	if len(remaining) > 0 {
		return ast.Program{}, fmt.Errorf("unexpected tokens after end of program: %v", remaining)
	}
	return ast.Program{
		Statements: statements,
	}, nil
}

func (p *Parser) parseStatements(tokens []tokenizer.Token) ([]ast.Statement, []tokenizer.Token) {
	statements := make([]ast.Statement, 0)
	remaining := tokens
	maxIterations := len(tokens) * 1000
	for len(remaining) > 0 {
		var stmt ast.Statement
		stmt, remaining = p.parseStatement(remaining)
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

func (p *Parser) parseStatement(tokens []tokenizer.Token) (ast.Statement, []tokenizer.Token) {
	sexp, remaining := p.parseSExp(tokens)
	if sexp != nil {
		return sexp, remaining
	}

	literal, remaining := p.parseLiteral(tokens)
	if literal != nil {
		return literal, remaining
	}
	return nil, remaining
}

func (p *Parser) parseSExp(tokens []tokenizer.Token) (ast.Statement, []tokenizer.Token) {
	tok, remaining := nextToken(tokens)
	if tok.Type != tokenizer.TokenTypeLParen {
		return nil, tokens
	}
	identifier, remaining := nextToken(remaining)
	if identifier.Type != tokenizer.TokenTypeIdentifier && identifier.Type != tokenizer.TokenTypeOperator {
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
		stmt, remaining = p.parseStatement(remaining)
		if stmt == nil {
			return nil, tokens
		}
		arguments = append(arguments, stmt)
	}
	if tok.Type != tokenizer.TokenTypeRParen {
		return nil, tokens
	}
	return ast.Function(identifier.Value, arguments), remaining
}

func (p *Parser) parseLiteral(tokens []tokenizer.Token) (ast.Statement, []tokenizer.Token) {
	tok, remaining := nextToken(tokens)
	switch tok.Type {
	case tokenizer.TokenTypeInteger:
		value, err := strconv.ParseInt(tok.Value, 10, 64)
		if err != nil {
			return nil, tokens
		}
		return ast.Integer(int(value)), remaining
	case tokenizer.TokenTypeFloat:
		value, err := strconv.ParseFloat(tok.Value, 64)
		if err != nil {
			return nil, tokens
		}
		return ast.Float(value), remaining
	case tokenizer.TokenTypeString:
		return ast.String(tok.Value), remaining
	default:
		return nil, tokens
	}
}

func nextToken(tokens []tokenizer.Token) (tokenizer.Token, []tokenizer.Token) {
	if len(tokens) == 0 {
		return tokenizer.EOF(), tokens
	}
	return tokens[0], tokens[1:]
}

func peekToken(tokens []tokenizer.Token) tokenizer.Token {
	if len(tokens) == 0 {
		return tokenizer.EOF()
	}
	return tokens[0]
}
