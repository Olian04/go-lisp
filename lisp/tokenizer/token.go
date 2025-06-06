package tokenizer

import "fmt"

type TokenType string

const (
	TokenTypeInteger    TokenType = "Integer"
	TokenTypeFloat      TokenType = "Float"
	TokenTypeString     TokenType = "String"
	TokenTypeIdentifier TokenType = "Identifier"
	TokenTypeLParen     TokenType = "LParen"
	TokenTypeRParen     TokenType = "RParen"
	TokenTypeEOF        TokenType = "EOF"
	TokenTypeInvalid    TokenType = "Invalid"
	TokenTypeNothing    TokenType = "Nothing"
	TokenTypeWhitespace TokenType = "Whitespace"
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
	case TokenTypeLParen:
		return "(LParen)"
	case TokenTypeRParen:
		return "(RParen)"
	case TokenTypeEOF:
		return "(EOF)"
	case TokenTypeNothing:
		return "(Nothing)"
	case TokenTypeInvalid:
		return fmt.Sprintf("(Invalid %s)", t.Value)
	case TokenTypeWhitespace:
		return fmt.Sprintf("(Whitespace %s)", t.Value)
	}
	return fmt.Sprintf("(Unknown %s)", t.Value)
}
