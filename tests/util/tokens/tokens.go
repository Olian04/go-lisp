package tokens

import "github.com/Olian04/go-lisp/lisp/tokenizer"

func Integer(value string) tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeInteger, Value: value}
}

func Float(value string) tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeFloat, Value: value}
}

func String(value string) tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeString, Value: value}
}

func Identifier(value string) tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeIdentifier, Value: value}
}

func LParen() tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeLParen, Value: "("}
}

func RParen() tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeRParen, Value: ")"}
}

func EOF() tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeEOF}
}

func Invalid(value string) tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeInvalid, Value: value}
}

func Whitespace(value string) tokenizer.Token {
	return tokenizer.Token{Type: tokenizer.TokenTypeWhitespace, Value: value}
}
