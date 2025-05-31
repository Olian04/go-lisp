package tokenizer

import "fmt"

type TokenType string

const (
	TokenTypeInteger    TokenType = "Integer"
	TokenTypeFloat      TokenType = "Float"
	TokenTypeString     TokenType = "String"
	TokenTypeIdentifier TokenType = "Identifier"
	TokenTypeOperator   TokenType = "Operator"
	TokenTypeLParen     TokenType = "LParen"
	TokenTypeRParen     TokenType = "RParen"
	TokenTypeEOF        TokenType = "EOF"
	TokenTypeInvalid    TokenType = "Invalid"
	tokenTypeWhitespace TokenType = "Whitespace"
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	switch t.Type {
	case TokenTypeInteger:
		return fmt.Sprintf("(Integer %s)", t.Value)
	case TokenTypeFloat:
		return fmt.Sprintf("(Float %s)", t.Value)
	case TokenTypeString:
		return fmt.Sprintf("(String %s)", t.Value)
	case TokenTypeIdentifier:
		return fmt.Sprintf("(Identifier %s)", t.Value)
	case TokenTypeOperator:
		return fmt.Sprintf("(Operator %s)", t.Value)
	case TokenTypeLParen:
		return "(LParen)"
	case TokenTypeRParen:
		return "(RParen)"
	case TokenTypeEOF:
		return "(EOF)"
	case TokenTypeInvalid:
		return fmt.Sprintf("(Invalid %s)", t.Value)
	case tokenTypeWhitespace:
		return fmt.Sprintf("(Whitespace %s)", t.Value)
	}
	return fmt.Sprintf("(Unknown %s)", t.Value)
}

func Integer(value string) Token {
	return Token{Type: TokenTypeInteger, Value: value}
}

func Float(value string) Token {
	return Token{Type: TokenTypeFloat, Value: value}
}

func String(value string) Token {
	return Token{Type: TokenTypeString, Value: value}
}

func Identifier(value string) Token {
	return Token{Type: TokenTypeIdentifier, Value: value}
}

func Operator(value string) Token {
	return Token{Type: TokenTypeOperator, Value: value}
}

func LParen() Token {
	return Token{Type: TokenTypeLParen, Value: "("}
}

func RParen() Token {
	return Token{Type: TokenTypeRParen, Value: ")"}
}

func EOF() Token {
	return Token{Type: TokenTypeEOF}
}

func Invalid(value string) Token {
	return Token{Type: TokenTypeInvalid, Value: value}
}

func Whitespace(value string) Token {
	return Token{Type: tokenTypeWhitespace, Value: value}
}
