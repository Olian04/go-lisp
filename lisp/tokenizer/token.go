package tokenizer

import "fmt"

type TokenType string

const (
	TokenTypeIdentifier TokenType = "Identifier"
	TokenTypeNumber     TokenType = "Number"
	TokenTypeString     TokenType = "String"
	TokenTypeBoolean    TokenType = "Boolean" // TODO: Implement this

	TokenTypeLParen   TokenType = "LParen"
	TokenTypeRParen   TokenType = "RParen"
	TokenTypeLBrace   TokenType = "LBrace"   // TODO: Implement this
	TokenTypeRBrace   TokenType = "RBrace"   // TODO: Implement this
	TokenTypeLBracket TokenType = "LBracket" // TODO: Implement this
	TokenTypeRBracket TokenType = "RBracket" // TODO: Implement this

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
	case TokenTypeNumber:
		return fmt.Sprintf("(Number %s)", t.Value)
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
